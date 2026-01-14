use num_format::{Buffer, Locale};

pub fn get_currency_symbol(currency_code: &str) -> &'static str {
    match currency_code.to_lowercase().as_str() {
        "usd" => "$",
        "eur" => "€",
        "gbp" => "£",
        "jpy" => "¥",
        "cny" => "¥",
        "krw" => "₩",
        "inr" => "₹",
        "rub" => "₽",
        "brl" => "R$",
        "cad" => "CA$",
        "aud" => "A$",
        "chf" => "CHF",
        "hkd" => "HK$",
        "sgd" => "S$",
        "sek" => "kr",
        "nok" => "kr",
        "dkk" => "kr",
        "pln" => "zł",
        "thb" => "฿",
        "idr" => "Rp",
        "myr" => "RM",
        "php" => "₱",
        "vnd" => "₫",
        "twd" => "NT$",
        "try" => "₺",
        "zar" => "R",
        "mxn" => "MX$",
        "nzd" => "NZ$",
        _ => "$",
    }
}

pub fn format_price_with_currency(price: f64, number_format: &str, currency: &str) -> String {
    let locale = match number_format {
        "european" => Locale::de, // German uses 1.234 (dot as thousand separator)
        _ => Locale::en,          // US/Asian uses 1,234 (comma as thousand separator)
    };

    let symbol = get_currency_symbol(currency);
    let rounded = price.round() as i64;

    let mut buf = Buffer::default();
    buf.write_formatted(&rounded, &locale);

    format!("{}{}", symbol, buf.as_str())
}

/// Format tray title from symbols with suffix
/// Example: FormatTrayTitle(["ETH", "BTC"], "$--,---") => "ETH $--,--- | BTC $--,---"
pub fn format_tray_title(symbols: &[String], suffix: &str) -> String {
    if symbols.is_empty() {
        return String::new();
    }

    symbols
        .iter()
        .map(|s| format!("{} {}", s, suffix))
        .collect::<Vec<_>>()
        .join(" | ")
}
