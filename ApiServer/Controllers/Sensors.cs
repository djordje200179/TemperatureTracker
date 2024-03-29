using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Security.Claims;
using System.Text.Json.Serialization;
using TemperatureTracker.Models;

namespace TemperatureTracker.ApiServer.Controllers;

[ApiController]
[Route("sensors")]
public class SensorsController(ReadingsContext readingsContext) : ControllerBase {
	[HttpGet]
	public async Task<ActionResult<IEnumerable<Sensor>>> GetAll() {
		return await readingsContext.Sensors.ToListAsync();
	}

	[HttpGet("{id:int}")]
	public async Task<ActionResult<Sensor>> Get(int id) {
		var sensor = await readingsContext.Sensors.FindAsync(id);
		if (sensor is null)
			return NotFound();

		return sensor;
	}

	[HttpGet("{id:int}/readings")]
	public async Task<ActionResult<IEnumerable<Reading>>> GetReadings(int id) {
		var sensor = await readingsContext.Sensors.FindAsync(id);
		if (sensor is null)
			return NotFound();

		await readingsContext.Sensors.Entry(sensor).Collection(s => s.Readings).LoadAsync();

		return Ok(sensor.Readings);
	}

	public record struct CreateSensorParams(
		[property: JsonPropertyName("name")] string Name,
		[property: JsonPropertyName("type")] string Type
	);

	[Authorize]
	[HttpPost]
	public async Task<ActionResult<Sensor>> Post([FromBody] CreateSensorParams sensorParams) {
		if (string.IsNullOrWhiteSpace(sensorParams.Name) || sensorParams.Name.Length > 30)
			return BadRequest("Invalid sensor name");

		if (string.IsNullOrWhiteSpace(sensorParams.Type) || sensorParams.Type.Length > 30)
			return BadRequest("Invalid sensor type");

		var claims = ((ClaimsIdentity)HttpContext.User.Identity!).Claims;

		var deviceName = claims.FirstOrDefault(claim => claim.Type == "device")?.Value;
		var device = await readingsContext.Devices.FirstOrDefaultAsync(d => d.Name == deviceName);
		if (device is null)
			return BadRequest("Device not found");

		if (await readingsContext.Sensors.AnyAsync(sensor => sensor.Name == sensorParams.Name && sensor.Device == device))
			return Conflict("Sensor name already exists");

		var sensor = new Sensor {
			Name = sensorParams.Name,
			Type = sensorParams.Type,
			Device = device
		};

		readingsContext.Sensors.Add(sensor);
		await readingsContext.SaveChangesAsync();

		return CreatedAtAction(nameof(Get), new { id = sensor.SensorId }, sensor);
	}
}
