fn convert_to_priority_value(character: &char) -> u32 {
    if !character.is_alphabetic() {
        panic!("Received character {} is not alphabetical", character);
    }

    let character_value: u32 = character.clone().into();
    if character.is_uppercase() {
        let initial_uppercase_value: u32 = 'A'.into();
        return character_value - initial_uppercase_value + 27;
    }
    let initial_lowercase_vaue: u32 = 'a'.into();
    return character_value - initial_lowercase_vaue + 1;
}
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_convert_to_priority_value_uppercase() {
        let character: char = 'a';
        let expected_value: u32 = 1;
        assert_eq!(expected_value, convert_to_priority_value(&character));
    }

    #[test]
    fn test_convert_to_priority_value_lowercase() {
        let character: char = 'A';
        let expected_value: u32 = 27;
        assert_eq!(expected_value, convert_to_priority_value(&character));
    }
}
