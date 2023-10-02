using System.ComponentModel.DataAnnotations;
using System.Net;
using System.Text.Json.Serialization;
using TemperatureTracker.Utils;

namespace TemperatureTracker.Models;

public class Device {
	[Key]
	[JsonPropertyName("id")]
	public int DeviceId { get; set; }

	[JsonPropertyName("address"), JsonConverter(typeof(IPAddressConverter))]
	public required IPAddress Address { get; set; }

	[MaxLength(30)]
	[JsonPropertyName("name")]
	public required string Name { get; set; }

	public ICollection<Sensor> Sensors { get; } = null!;
}