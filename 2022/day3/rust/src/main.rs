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
