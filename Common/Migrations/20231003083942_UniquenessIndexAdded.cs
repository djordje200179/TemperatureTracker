﻿using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace TemperatureTracker.Migrations
{
    /// <inheritdoc />
    public partial class UniquenessIndexAdded : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateIndex(
                name: "IX_Devices_Name",
                table: "Devices",
                column: "Name",
                unique: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropIndex(
                name: "IX_Devices_Name",
                table: "Devices");
        }
    }
}
