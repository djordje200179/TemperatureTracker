using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace TemperatureTracker.Models;

public class Reading {
	[Key]
	[JsonPropertyName("id")]
	public int ReadingId { get; set; }

	[JsonPropertyName("time")]
	public DateTime Time { get; set; } = DateTime.Now;

	[JsonPropertyName("temperature")]
	public double? Temperature { get; set; }

	[JsonPropertyName("humidity")]
	public double? Humidity { get; set; }

	[JsonPropertyName("sensor")]
	public int SensorId { get; set; }

	[JsonIgnore]
	public Sensor Sensor { get; set; } = null!;
}