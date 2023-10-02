using TemperatureTracker;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddDbContext<ReadingsContext>();

var app = builder.Build();

app.MapControllers();

app.Run();