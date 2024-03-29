﻿// <auto-generated />
using System;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Metadata;
using Microsoft.EntityFrameworkCore.Migrations;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using TemperatureTracker;

#nullable disable

namespace TemperatureTracker.Migrations
{
    [DbContext(typeof(ReadingsContext))]
    [Migration("20231003083942_UniquenessIndexAdded")]
    partial class UniquenessIndexAdded
    {
        /// <inheritdoc />
        protected override void BuildTargetModel(ModelBuilder modelBuilder)
        {
#pragma warning disable 612, 618
            modelBuilder
                .HasAnnotation("ProductVersion", "7.0.11")
                .HasAnnotation("Relational:MaxIdentifierLength", 128);

            SqlServerModelBuilderExtensions.UseIdentityColumns(modelBuilder);

            modelBuilder.Entity("TemperatureTracker.Models.Device", b =>
                {
                    b.Property<int>("DeviceId")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("int")
                        .HasAnnotation("Relational:JsonPropertyName", "id");

                    SqlServerPropertyBuilderExtensions.UseIdentityColumn(b.Property<int>("DeviceId"));

                    b.Property<string>("Address")
                        .IsRequired()
                        .HasColumnType("nvarchar(45)")
                        .HasAnnotation("Relational:JsonPropertyName", "address");

                    b.Property<string>("Key")
                        .IsRequired()
                        .HasMaxLength(30)
                        .HasColumnType("nvarchar(30)");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasMaxLength(30)
                        .HasColumnType("nvarchar(30)")
                        .HasAnnotation("Relational:JsonPropertyName", "name");

                    b.HasKey("DeviceId");

                    b.HasIndex("Name")
                        .IsUnique();

                    b.ToTable("Devices");
                });

            modelBuilder.Entity("TemperatureTracker.Models.Reading", b =>
                {
                    b.Property<int>("ReadingId")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("int")
                        .HasAnnotation("Relational:JsonPropertyName", "id");

                    SqlServerPropertyBuilderExtensions.UseIdentityColumn(b.Property<int>("ReadingId"));

                    b.Property<double?>("Humidity")
                        .HasColumnType("float")
                        .HasAnnotation("Relational:JsonPropertyName", "humidity");

                    b.Property<int>("SensorId")
                        .HasColumnType("int")
                        .HasAnnotation("Relational:JsonPropertyName", "sensor");

                    b.Property<double?>("Temperature")
                        .HasColumnType("float")
                        .HasAnnotation("Relational:JsonPropertyName", "temperature");

                    b.Property<DateTime>("Time")
                        .HasColumnType("datetime2")
                        .HasAnnotation("Relational:JsonPropertyName", "time");

                    b.HasKey("ReadingId");

                    b.HasIndex("SensorId");

                    b.ToTable("Readings");
                });

            modelBuilder.Entity("TemperatureTracker.Models.Sensor", b =>
                {
                    b.Property<int>("SensorId")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("int")
                        .HasAnnotation("Relational:JsonPropertyName", "id");

                    SqlServerPropertyBuilderExtensions.UseIdentityColumn(b.Property<int>("SensorId"));

                    b.Property<int>("DeviceId")
                        .HasColumnType("int")
                        .HasAnnotation("Relational:JsonPropertyName", "device");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasMaxLength(30)
                        .HasColumnType("nvarchar(30)")
                        .HasAnnotation("Relational:JsonPropertyName", "name");

                    b.HasKey("SensorId");

                    b.HasIndex("DeviceId");

                    b.ToTable("Sensors");
                });

            modelBuilder.Entity("TemperatureTracker.Models.Reading", b =>
                {
                    b.HasOne("TemperatureTracker.Models.Sensor", "Sensor")
                        .WithMany("Readings")
                        .HasForeignKey("SensorId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Sensor");
                });

            modelBuilder.Entity("TemperatureTracker.Models.Sensor", b =>
                {
                    b.HasOne("TemperatureTracker.Models.Device", "Device")
                        .WithMany("Sensors")
                        .HasForeignKey("DeviceId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Device");
                });

            modelBuilder.Entity("TemperatureTracker.Models.Device", b =>
                {
                    b.Navigation("Sensors");
                });

            modelBuilder.Entity("TemperatureTracker.Models.Sensor", b =>
                {
                    b.Navigation("Readings");
                });
#pragma warning restore 612, 618
        }
    }
}
