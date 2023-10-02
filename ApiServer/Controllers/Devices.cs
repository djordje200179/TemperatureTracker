using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Text.Json.Serialization;
using TemperatureTracker.Models;

namespace TemperatureTracker.ApiServer.Controllers;

[ApiController]
[Route("devices")]
public class DevicesController : ControllerBase {
	private readonly ReadingsContext readingsContext;

	public DevicesController(ReadingsContext readingsContext) {
		this.readingsContext = readingsContext;
	}

	[HttpGet]
	public async Task<IEnumerable<Device>> GetAll() {
		return await readingsContext.Devices.ToListAsync();
	}

	[HttpGet("{id}")]
	public async Task<ActionResult<Device>> Get(int id) {
		var device = await readingsContext.Devices.FindAsync(id);
		if (device is null)
			return NotFound();

		return device;
	}

	public record struct CreateDeviceParams(
		[property: JsonPropertyName("name")] string Name
	);

	[HttpPost]
	public async Task<ActionResult<Device>> Post([FromBody] CreateDeviceParams deviceParams) {
		if (string.IsNullOrWhiteSpace(deviceParams.Name) || deviceParams.Name.Length > 30)
			return BadRequest("Invalid device name");

		var address = HttpContext.Connection.RemoteIpAddress;

		if (address is null)
			return BadRequest("Invalid IP address");

		address = address.MapToIPv4();

		var device = new Device {
			Address = address,
			Name = deviceParams.Name,
		};

		readingsContext.Devices.Add(device);
		await readingsContext.SaveChangesAsync();

		return CreatedAtAction(nameof(Get), new { id = device.DeviceId }, device);
	}
}