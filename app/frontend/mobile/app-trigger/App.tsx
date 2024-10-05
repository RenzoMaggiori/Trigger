import React, { useState, useEffect } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import LandingPage from './app/LandingPage';
import SignIn from './app/SignIn';
import SignUp from './app/SignUp';
import TabLayout from './navigation/TabLayout';
import Settings from './app/Settings';

const Stack = createStackNavigator();

export default function App() {
    const [isSignedIn, setIsSignedIn] = useState(false);

    useEffect(() => {
        const checkAuthStatus = async () => {
            const userAuthenticated = false; // TODO: Implement authentication check
            setIsSignedIn(userAuthenticated);
        };
        checkAuthStatus();
    }, []);

    return (
        <NavigationContainer>
            <Stack.Navigator>
                {isSignedIn ? (
                    <Stack.Screen
                        name="Tabs"
                        component={TabLayout}
                        options={{ headerShown: false }}
                    />
                ) : (
                    <>
                        <Stack.Screen
                            name="LandingPage"
                            component={LandingPage}
                            options={{ headerShown: false }}
                        />
                        <Stack.Screen
                            name="SignIn"
                            component={SignIn}
                            options={{ headerShown: false }}
                        />
                        <Stack.Screen
                            name="SignUp"
                            component={SignUp}
                            options={{ headerShown: false }}
                        />
                        <Stack.Screen
                            name="Settings"
                            component={Settings}
                            options={{ headerShown: false }}
                        />
                    </>
                )}
            </Stack.Navigator>
        </NavigationContainer>
    );
}
