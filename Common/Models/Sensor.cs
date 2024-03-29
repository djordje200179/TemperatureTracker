﻿using Microsoft.EntityFrameworkCore;
using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace TemperatureTracker.Models;

[Index(nameof(DeviceId), nameof(Name), IsUnique = true)]
public class Sensor {
	[Key]
	[JsonPropertyName("id")]
	public int SensorId { get; set; }

	[MaxLength(30)]
	[JsonPropertyName("name")]
	public required string Name { get; set; }

	[MaxLength(30)]
	[JsonPropertyName("type")]
	public required string Type { get; set; }

	[JsonPropertyName("device")]
	public int DeviceId { get; set; }

	[JsonIgnore]
	public Device Device { get; set; } = null!;

	[JsonIgnore]
	public ICollection<Reading> Readings { get; } = null!;
}