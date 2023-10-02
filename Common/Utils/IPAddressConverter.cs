using System.Net;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace TemperatureTracker.Utils;

public class IPAddressConverter : JsonConverter<IPAddress> {
	public override IPAddress? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options) {
		var address = reader.GetString();
		if (address is null)
			return null;

		return IPAddress.Parse(address);
	}

	public override void Write(Utf8JsonWriter writer, IPAddress value, JsonSerializerOptions options) {
		writer.WriteStringValue(value.ToString());
	}
}