import pydantic

from its_a_wonderful_world.typings import WebsocketActionsTypes


class WebsocketMessageData(pydantic.BaseModel):
    pass


class WebsocketMessage(pydantic.BaseModel):
    action: WebsocketActionsTypes
    data: WebsocketMessageData


class GameCardItem(pydantic.BaseModel):
    name: str
    image: str
    card_type: str
    category: str
    image: str
    recycling_bonus: str
    has_construction_bonus: bool = pydantic.Field(default=False)
    has_character_token_bonus: bool = pydantic.Field(default=False)

    # @pydantic.field_validator('recycling_bonus')
    # @classmethod
    # def set_has_construction_bonus(cls, v, values):
    #     return values.get('has_construction_bonus', False)


class GameCardsResponse(pydantic.BaseModel):
    data: list[GameCardItem] = []
