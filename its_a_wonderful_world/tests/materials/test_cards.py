import unittest

from its_a_wonderful_world.material.card import BaseDevelopmentCardMixin
from its_a_wonderful_world.typings import DevelopmentTypes, ResourceCubesTypes


class TestCard(unittest.TestCase):
    def test_instance(self):
        instance = BaseDevelopmentCardMixin()
        self.assertTrue(
            instance.recycling_bonus == ResourceCubesTypes.MATERIAL
        )
        self.assertFalse(instance.is_face_down)
        self.assertEqual(instance.number_of_copies(), 0)
        self.assertTrue(instance.construction_cost_count == 0)
        self.assertTrue(instance.construction_cost_resource_types_count == 0)

    def test_repr(self):
        instance = BaseDevelopmentCardMixin(name="Test Card")
        repr_str = repr(instance)
        self.assertIn("Test Card", repr_str)
        self.assertIn("BaseDevelopmentCardMixin", repr_str)

    def test_hash(self):
        instance1 = BaseDevelopmentCardMixin()
        instance2 = BaseDevelopmentCardMixin()
        self.assertNotEqual(hash(instance1), hash(instance2))

    def test_add_resource(self):
        instance = BaseDevelopmentCardMixin(name='Basic Card')

        # Adding to a card that has not resource
        self.assertRaises(
            ValueError,
            instance.add_resource,
            ResourceCubesTypes.MATERIAL
        )

        instance.construction_cost = [
            ResourceCubesTypes.MATERIAL
        ]

        # Adding to a face-down card
        instance.is_face_down = True
        self.assertRaises(
            ValueError,
            instance.add_resource,
            ResourceCubesTypes.MATERIAL
        )

        # Adding to a completed card
        instance.is_face_down = False
        instance.construction_progress = [
            ResourceCubesTypes.MATERIAL
        ]
        self.assertRaises(
            ValueError,
            instance.add_resource,
            ResourceCubesTypes.MATERIAL
        )

        # Adding to a card normally
        instance.construction_progress = []
        instance.add_resource(ResourceCubesTypes.MATERIAL)
        self.assertIn(
            ResourceCubesTypes.MATERIAL,
            instance.construction_progress
        )
        self.assertEqual(
            instance.construction_progress_percentage,
            1.0
        )

    def test_construction_bonus_count(self):
        cards = [
            BaseDevelopmentCardMixin(
                name='Card 1',
                card_type=DevelopmentTypes.VEHICLE
            ),
            BaseDevelopmentCardMixin(
                name='Card 2',
                card_type=DevelopmentTypes.VEHICLE
            )
        ]

        instance = BaseDevelopmentCardMixin(
            name='Bonus Card',
            multiply_construction_bonus_constraint=DevelopmentTypes.VEHICLE
        )

        count = instance.construction_bonus_count(cards)
        self.assertEqual(count, 3)  # 2 matching cards + 1 for itself
