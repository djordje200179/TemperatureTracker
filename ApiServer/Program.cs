using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.IdentityModel.Tokens;
using System.Text;
using TemperatureTracker;

var builder = WebApplication.CreateBuilder(args);

builder.Services
	.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
	.AddJwtBearer(options => {
		options.TokenValidationParameters = new TokenValidationParameters {
			ValidateLifetime = true,

			ValidateIssuer = true,
			ValidIssuer = builder.Configuration["Jwt:Issuer"]!,

			ValidateAudience = true,
			ValidAudience = builder.Configuration["Jwt:Audience"]!,

			ValidateIssuerSigningKey = true,
			IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(builder.Configuration["Jwt:Key"]!)),
		};
	});

builder.Services.AddAuthorization();

builder.Services.AddControllers();
builder.Services.AddDbContext<ReadingsContext>();

var app = builder.Build();

app.UseAuthentication().UseAuthorization();

app.MapControllers();

app.Run();