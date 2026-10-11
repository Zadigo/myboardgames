enum CardType {
    Standard,
    StandardSpecialCard
}

#[derive(Debug,PartialEq)]
enum CardOperator {
    Add,
    Subtract,
    Multiply,
    Divide,
    Freeze ,
    Flip3,
    SecondChance,
}

struct BaseCard {
    uuid: String,
    card_type: CardType,
    card_value: i32,
    card_operator: CardOperator,
    discarded: bool,
}

trait BaseCardTrait{
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
    fn new(card_type: CardType) -> BaseCard {
        Self {
            uuid: String::new(),
            card_operator: CardOperator::Add,
            card_type,
            card_value: 0,
            discarded: false
        }
    }
}

impl BaseCardTrait for BaseCard {
    fn get_uuid(self) -> String {
        self.uuid
    }

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

    #[test]
    fn test_base_card_instance() {
        let card = BaseCard::new(CardType::Standard);
    }
}
