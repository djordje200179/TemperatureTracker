using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Text.Json.Serialization;
using TemperatureTracker.Models;

namespace TemperatureTracker.ApiServer.Controllers;

[ApiController]
[Route("readings")]
public class ReadingsController : ControllerBase {
	private readonly ReadingsContext readingsContext;

	public ReadingsController(ReadingsContext readingsContext) {
		this.readingsContext = readingsContext;
	}

	[HttpGet]
	public async Task<ActionResult<IEnumerable<Reading>>> GetAll() {
		return await readingsContext.Readings.ToListAsync();
	}

	[HttpGet("{id}")]
	public async Task<ActionResult<Reading>> Get(int id) {
		var reading = await readingsContext.Readings.FindAsync(id);
		if (reading == null)
			return NotFound();

		return reading;
	}

	public record struct CreateReadingParams(
		[property: JsonPropertyName("sensor")] int Sensor,
		[property: JsonPropertyName("temperature")] double? Temperature,
		[property: JsonPropertyName("humidity")] double? Humidity
	);

	[Authorize]
	[HttpPost]
	public async Task<ActionResult<Reading>> Create(CreateReadingParams readingParams) {
		var sensor = await readingsContext.Sensors.FindAsync(readingParams.Sensor);
		if (sensor is null)
			return BadRequest("Sensor not found");

		var reading = new Reading {
			Sensor = sensor,
			Temperature = readingParams.Temperature,
			Humidity = readingParams.Humidity
		};

		readingsContext.Readings.Add(reading);
		await readingsContext.SaveChangesAsync();

		return CreatedAtAction(nameof(Get), new { id = reading.ReadingId }, reading);
	}
}
