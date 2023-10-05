using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using System.Text.Json.Serialization;
using TemperatureTracker.Models;

namespace TemperatureTracker.ApiServer.Controllers;

[ApiController]
[Route("devices")]
public class DevicesController : ControllerBase {
	private readonly ReadingsContext readingsContext;
	private readonly IConfiguration configuration;

	public DevicesController(ReadingsContext readingsContext, IConfiguration configuration) {
		this.readingsContext = readingsContext;
		this.configuration = configuration;
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
		[property: JsonPropertyName("name")] string Name,
		[property: JsonPropertyName("key")] string Key
	);

	[HttpPost]
	public async Task<ActionResult<Device>> Post([FromBody] CreateDeviceParams deviceParams) {
		if (string.IsNullOrWhiteSpace(deviceParams.Name) || deviceParams.Name.Length > 30)
			return BadRequest("Invalid device name");

		if (string.IsNullOrWhiteSpace(deviceParams.Key) || deviceParams.Key.Length > 30)
			return BadRequest("Invalid device key");

		if (await readingsContext.Devices.AnyAsync(d => d.Name == deviceParams.Name))
			return Conflict("Device name already exists");

		var address = HttpContext.Connection.RemoteIpAddress;

		if (address is null)
			return BadRequest("Invalid IP address");

		address = address.MapToIPv4();

		var device = new Device {
			Address = address,
			Name = deviceParams.Name,
			Key = deviceParams.Key
		};

		readingsContext.Devices.Add(device);
		await readingsContext.SaveChangesAsync();

		return CreatedAtAction(nameof(Get), new { id = device.DeviceId }, device);
	}

	public record struct DeviceAuthParams(
		[property: JsonPropertyName("name")] string Name,
		[property: JsonPropertyName("key")] string Key
	);

	[HttpPost]
	[Route("auth")]
	public async Task<ActionResult<string>> Authenticate([FromBody] DeviceAuthParams authParams) {
		var device = await readingsContext.Devices.FirstOrDefaultAsync(d => d.Name == authParams.Name);
		if (device is null)
			return BadRequest("Device not found");

		var claims = new List<Claim> {
			new (JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
			new (JwtRegisteredClaimNames.Sub, authParams.Name),
			new ("device", authParams.Name)
		};	

		var tokenDuration = TimeSpan.FromDays(1);
		var tokenKey = configuration["Jwt:Key"]!;
		var tokenIssuer = configuration["Jwt:Issuer"]!;
		var tokenAudience = configuration["Jwt:Audience"]!;

		var tokenDescriptor = new SecurityTokenDescriptor {
			Subject = new ClaimsIdentity(claims),
			Issuer = tokenIssuer,
			Audience = tokenAudience,
			Expires = DateTime.UtcNow + tokenDuration,
			SigningCredentials = new SigningCredentials(
				new SymmetricSecurityKey(Encoding.UTF8.GetBytes(tokenKey)), SecurityAlgorithms.HmacSha256Signature
			)
		};

		var tokenHandler = new JwtSecurityTokenHandler();
		var token = tokenHandler.CreateToken(tokenDescriptor);
		var jwt = tokenHandler.WriteToken(token);

		return jwt;
	}
}