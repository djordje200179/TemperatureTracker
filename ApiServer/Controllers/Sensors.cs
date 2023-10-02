using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Text.Json.Serialization;
using TemperatureTracker.Models;

namespace TemperatureTracker.ApiServer.Controllers;

[ApiController]
[Route("sensors")]
public class SensorsController : ControllerBase {
	private readonly ReadingsContext readingsContext;

	public SensorsController(ReadingsContext readingsContext) {
		this.readingsContext = readingsContext;
	}

	[HttpGet]
	public async Task<ActionResult<IEnumerable<Sensor>>> GetAll() {
		return await readingsContext.Sensors.ToListAsync();
	}

	[HttpGet("{id}")]
	public async Task<ActionResult<Sensor>> Get(int id) {
		var sensor = await readingsContext.Sensors.FindAsync(id);
		if (sensor is null)
			return NotFound();

		return sensor;
	}

	[HttpGet("{id}/readings")]
	public async Task<ActionResult<IEnumerable<Reading>>> GetReadings(int id) {
		var sensor = await readingsContext.Sensors.FindAsync(id);
		if (sensor is null)
			return NotFound();

		readingsContext.Sensors.Entry(sensor).Collection(s => s.Readings).Load();

		return Ok(sensor.Readings);
	}

	public record struct CreateSensorParams(
		[property: JsonPropertyName("deviceName")] string DeviceName,
		[property: JsonPropertyName("sensorName")] string SensorName
	);

	[HttpPost]
	public async Task<ActionResult<Sensor>> Post([FromBody] CreateSensorParams sensorParams) {
		if (string.IsNullOrWhiteSpace(sensorParams.DeviceName) || sensorParams.DeviceName.Length > 30)
			return BadRequest("Invalid device name");

		if (string.IsNullOrWhiteSpace(sensorParams.SensorName) || sensorParams.SensorName.Length > 30)
			return BadRequest("Invalid sensor name");

		var device = await readingsContext.Devices.FirstOrDefaultAsync(d => d.Name == sensorParams.DeviceName);
		if (device is null)
			return BadRequest("Device not found");

		var sensor = new Sensor {
			Name = sensorParams.SensorName,
			Device = device
		};

		readingsContext.Sensors.Add(sensor);
		await readingsContext.SaveChangesAsync();

		return CreatedAtAction(nameof(Get), new { id = sensor.SensorId }, sensor);
	}
}
