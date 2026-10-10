use std::str::FromStr;

use clap::Parser;
use clap_derive::{Parser, ValueEnum};

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(long)]
    temperature: Option<i32>,

    #[arg(long)]
    convert_to: Option<String>
}

#[derive(ValueEnum, Clone, Debug, PartialEq)]
enum TemperatureUnit {
    Celsius,
    Kelvin,
    Fahrenheit
}

// Implement FromStr so we can convert "celsius" (String) into TemperatureUnit::Celsius
impl FromStr for TemperatureUnit {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.trim().to_lowercase().as_str() {
            "celsius" | "c" => Ok(TemperatureUnit::Celsius),
            "kelvin" | "k" => Ok(TemperatureUnit::Kelvin),
            "fahrenheit" | "f" => Ok(TemperatureUnit::Fahrenheit),
            _ => Err(()),
        }
    }
}


fn celsius_to_farenheit(v: i32) -> i32 {
    (v * 9/5) + 32
}

pub fn temperature_exercise() {
    let args = Args::parse();
    let unit = args.convert_to
        .as_deref()
        .and_then(|s| s.parse::<TemperatureUnit>().ok());

    if unit == Some(TemperatureUnit::Celsius) {
        let calc_result = celsius_to_farenheit(args.temperature.unwrap());
        print!("{:?}°C is {}°F", unit, calc_result)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_celsius_to_farenheit() {
        let result = celsius_to_farenheit(15);
        assert_eq!(result > 0, true)
    }
}
