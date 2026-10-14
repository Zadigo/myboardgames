#[derive(Debug,PartialEq)]
pub enum CardType {
    Standard,
    StandardSpecialCard
}

#[derive(Debug,PartialEq)]
pub enum CardOperator {
    Add,
    Subtract,
    Multiply,
    Divide,
    Freeze ,
    Flip3,
    SecondChance,
}

#[derive(Debug, PartialEq)]
pub struct BaseCard {
    pub uuid: String,
    pub card_type: CardType,
    pub card_value: i32,
    pub card_operator: CardOperator,
    pub discarded: bool,
}

trait BaseCardTrait {
    fn new() -> BaseCard;
    fn get_uuid(self) -> String;
    fn resolve(self) -> i32;
    fn set_player(&mut self, uuid: String);
    fn get_player(self) -> String;
    fn get_value(self) -> i32;
    fn is_special(self) -> bool;
    fn get_operator(&self) -> &str;
    fn is_card_operation(self) -> bool;
}

impl BaseCard {
    pub fn new(card_type: Option<CardType>, card_operator: Option<CardOperator>) -> BaseCard {
        Self {
            uuid: String::new(),
            card_operator: card_operator.unwrap_or(CardOperator::Add),
            card_type: card_type.unwrap_or(CardType::Standard),
            card_value: 0,
            discarded: false
        }
    }
}

impl BaseCardTrait for BaseCard {
    fn new() -> BaseCard {
        BaseCard {
            uuid: String::from("1345"),
            card_operator: CardOperator::Add,
            card_type: CardType::Standard,
            card_value: 0,
            discarded: false
        }
    }


    fn get_uuid(self) -> String {
        self.uuid
    }

    // Function used to resolve the card based on
    // the options provided by the user
    fn resolve(self) -> i32 {
        0
    }

    fn get_operator(&self) -> &str {
        match self.card_operator {
            CardOperator::Add => "+",
            CardOperator::Subtract => "-",
            CardOperator::Multiply => "*",
            CardOperator::Divide => "/",
            CardOperator::Freeze => "F",
            CardOperator::Flip3=> "3",
            CardOperator::SecondChance => "S",
        }
    }

    fn get_player(self) -> String {
        self.uuid
    }

    fn get_value(self) -> i32 {
        self.card_value
    }

    fn is_card_operation(self) -> bool {
        let special_operators = [
            CardOperator::Add,
            CardOperator::Divide,
            CardOperator::Divide,
            CardOperator::Multiply,
        ];

        special_operators.contains(&self.card_operator)
    }

    fn is_special(self) -> bool {
        match self.card_operator {
            CardOperator::Flip3 => true,
            CardOperator::SecondChance => true,
            CardOperator::Freeze => true,
            _ => false,
        }
    }

    fn set_player(&mut self, uuid: String) {
        self.uuid = uuid
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    // use assertables::*; 

    #[test]
    fn test_base_standard_card_instance() {
        let card_types = [
            CardType::Standard,
            CardType::StandardSpecialCard
        ];

        for v in card_types {
            let card = BaseCard::new(Some(v), Some(CardOperator::Add));
            assert_eq!(card.card_type, CardType::Standard);
            assert_eq!(card.card_value, 0)
        }
    }

    // #[test]
    // fn test_get_operator() {
    //     let operators = [
    //         CardOperator::Add,
    //         CardOperator::Divide,
    //         CardOperator::Multiply,
    //         CardOperator::Flip3,
    //         CardOperator::Freeze,
    //         CardOperator::SecondChance,
    //         CardOperator::Subtract,
    //     ];

    //     for v in operators.into_iter() {
    //         let card = BaseCard::new(Some(CardType::Standard), Some(v));
    //         assert_eq!(card.card_operator, v)
    //     }
    // }
}
