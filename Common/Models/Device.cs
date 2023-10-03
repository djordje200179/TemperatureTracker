using Microsoft.EntityFrameworkCore;
using System.ComponentModel.DataAnnotations;
using System.Net;
using System.Text.Json.Serialization;
using TemperatureTracker.Utils;

namespace TemperatureTracker.Models;

[Index(nameof(Name), IsUnique = true)]
public class Device {
	[Key]
	[JsonPropertyName("id")]
	public int DeviceId { get; set; }

	[JsonPropertyName("address"), JsonConverter(typeof(IPAddressConverter))]
	public required IPAddress Address { get; set; }

	[MaxLength(30)]
	[JsonPropertyName("name")]
	public required string Name { get; set; }

	[MaxLength(30)]
	[JsonIgnore]
	public string Key { get; set; } = null!;

	public ICollection<Sensor> Sensors { get; } = null!;
}