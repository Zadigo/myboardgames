use crate::game::card::{BaseCard, CardOperator, CardType};

pub fn standard_card_factory() -> Vec<BaseCard> {
    let mut cards = Vec::new();
    for i in 1..=12 {
        let mut card = BaseCard::new(
            Some(CardType::Standard),
            Some(CardOperator::Add)
        );
        card.card_value = i;
        cards.push(card);
    }
    
    cards
}

pub fn standard_special_cards_factory() -> Vec<BaseCard> {
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

# [cfg(test)]
mod tests {
    use super::*;

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
}
