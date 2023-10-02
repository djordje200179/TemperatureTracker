using Microsoft.EntityFrameworkCore;
using TemperatureTracker.Models;

namespace TemperatureTracker;

public class ReadingsContext : DbContext {
	public DbSet<Device> Devices { get; set; }
	public DbSet<Sensor> Sensors { get; set; }
	public DbSet<Reading> Readings { get; set; }

	protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder) {
		optionsBuilder.UseSqlServer("Data Source=localhost;Persist Security Info=True;User ID=sa;Password=Djole2001;TrustServerCertificate=True");
	}

	protected override void OnModelCreating(ModelBuilder modelBuilder) {
		modelBuilder.Entity<Device>().Navigation(device => device.Sensors).AutoInclude();

		modelBuilder.Entity<Sensor>().Navigation(sensor => sensor.Device).AutoInclude();

		modelBuilder.Entity<Reading>().Navigation(reading => reading.Sensor).AutoInclude();
	}
}