using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Security.Claims;
using System.Text.Json.Serialization;
using TemperatureTracker.Models;

namespace TemperatureTracker.ApiServer.Controllers;

[ApiController]
[Route("readings")]
public class ReadingsController(ReadingsContext readingsContext) : ControllerBase {
	[HttpGet]
	public async Task<ActionResult<IEnumerable<Reading>>> GetAll() {
		return await readingsContext.Readings.ToListAsync();
	}

	[HttpGet("{id:int}")]
	public async Task<ActionResult<Reading>> Get(int id) {
		var reading = await readingsContext.Readings.FindAsync(id);
		if (reading is null)
			return NotFound();

		return reading;
	}

	public record struct CreateReadingParams(
		[property: JsonPropertyName("sensor")] string SensorName,
		[property: JsonPropertyName("temperature")] double? Temperature,
		[property: JsonPropertyName("humidity")] double? Humidity
	);

	[Authorize]
	[HttpPost]
	public async Task<ActionResult<Reading>> Create(CreateReadingParams readingParams) {
		var claims = ((ClaimsIdentity)HttpContext.User.Identity!).Claims;

		var deviceName = claims.FirstOrDefault(claim => claim.Type == "device")?.Value;
		var sensor = await readingsContext.Sensors.FirstOrDefaultAsync(
			sensor => sensor.Name == readingParams.SensorName && sensor.Device.Name == deviceName
		);
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
