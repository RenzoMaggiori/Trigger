import React from 'react';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { MaterialIcons, MaterialCommunityIcons } from '@expo/vector-icons';
import HomeScreen from '../app/(tabs)/HomeScreen';
import TriggersScreen from '../app/(tabs)/TriggersScreen';
import TemplatesScreen from '../app/(tabs)/TemplatesScreen';

const Tab = createBottomTabNavigator();

export default function TabLayout() {
    return (
        <Tab.Navigator screenOptions={{ headerShown: false }}>
            <Tab.Screen
                name="Home"
                component={HomeScreen}
                options={{
                    tabBarIcon: ({ color, size }) => (
                        <MaterialIcons name="home" color={color} size={size} />
                    )
                }}
            />
            <Tab.Screen
                name="Triggers"
                component={TriggersScreen}
                options={{
                    tabBarIcon: ({ color, size }) => (
                        <MaterialIcons name="link" color={color} size={size} />
                    )
                }}
            />
            <Tab.Screen
                name="Templates"
                component={TemplatesScreen}
                options={{
                    tabBarIcon: ({ color, size }) => (
                        <MaterialCommunityIcons name="file-tree" color={color} size={size} />
                    )
                }}
            />
        </Tab.Navigator>
    );
}
