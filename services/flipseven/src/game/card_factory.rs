use crate::game::card::{BaseCard, CardOperator, CardType};

fn standard_card_factory() -> Vec<BaseCard> {
    let mut cards = Vec::with_capacity(70);

    for i in 0..=12i32 {
        let count = if i == 0 { 1 } else { i };
        let mut card = BaseCard::new(Some(CardType::Standard), Some(CardOperator::Add));

        card.card_value = i;
        cards.extend(std::iter::repeat_n(card, count as usize));
    }

    cards
}

fn standard_special_cards_factory() -> Vec<BaseCard> {
    let card_types = [
        CardOperator::Flip3,
        CardOperator::SecondChance,
        CardOperator::Freeze
    ];

    let mut cards = Vec::new();

    for _ in card_types.iter() {
        let card = BaseCard::new(Some(CardType::StandardSpecialCard), None);
        cards.push(card);
    }

    cards
}

pub fn create_standard_game_cards() -> Vec<BaseCard> {
    let standard_cards = standard_card_factory();
    let special_cards = standard_special_cards_factory();
    standard_cards.into_iter().chain(special_cards).collect()
}

# [cfg(test)]
mod tests {
    use crate::game::card::BaseCardTrait;

use super::*;
    use assertables::*;

    #[test]
    fn test_standard_card_factory() {
        let cards = standard_card_factory();
        assert_eq!(cards.len(), 12);
    }

    #[test]
    fn test_standard_special_cards_factory() {
        let cards = standard_special_cards_factory();
        assert_eq!(cards.len(), 3);
    }

    #[test]
    fn test_create_all_game_cards() {
        let cards = create_standard_game_cards();
        assert_eq!(cards.len(), 82);
    }

    fn test_has_special_cards() {
        let cards = create_standard_game_cards();
        let filtered_cards = cards.into_iter().filter(|c| c.clone().is_special());
        assert_eq!(filtered_cards.count(), 5);
    }
}
