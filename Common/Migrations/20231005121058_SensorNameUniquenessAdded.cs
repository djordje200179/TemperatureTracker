using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace TemperatureTracker.Migrations
{
    /// <inheritdoc />
    public partial class SensorNameUniquenessAdded : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropIndex(
                name: "IX_Sensors_DeviceId",
                table: "Sensors");

            migrationBuilder.CreateIndex(
                name: "IX_Sensors_DeviceId_Name",
                table: "Sensors",
                columns: new[] { "DeviceId", "Name" },
                unique: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropIndex(
                name: "IX_Sensors_DeviceId_Name",
                table: "Sensors");

            migrationBuilder.CreateIndex(
                name: "IX_Sensors_DeviceId",
                table: "Sensors",
                column: "DeviceId");
        }
    }
}
