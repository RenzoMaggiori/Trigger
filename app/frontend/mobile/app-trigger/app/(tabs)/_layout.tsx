import React, { useState } from 'react';
import { View, Image, TouchableOpacity, StyleSheet } from 'react-native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { MaterialIcons } from '@expo/vector-icons';
import HomeScreen from './HomeScreen';
import TriggersScreen from './TriggersScreen';
import CommunityScreen from './CommunityScreen';
import { Menu, Divider, Provider } from 'react-native-paper';
import { useRouter } from 'expo-router';
import { Colors } from '../../constants/Colors';

const Tab = createBottomTabNavigator();

export default function TabLayout() {
  const [menuVisible, setMenuVisible] = useState(false);

  const openMenu = () => setMenuVisible(true);
  const closeMenu = () => setMenuVisible(false);

  const router = useRouter();

  const handleLogout = () => {
    router.push('/LandingPage');
  }

  return (
    <Provider>
      <View style={styles.headerContainer}>
        <Image source={require('../../assets/images/react-logo.png')} style={styles.logo} />

        <Menu
          visible={menuVisible}
          onDismiss={closeMenu}
          anchor={
            <TouchableOpacity onPress={openMenu}>
              <MaterialIcons name="account-circle" size={32} color={Colors.light.text} />
            </TouchableOpacity>
          }
        >
          <Menu.Item onPress={() => { }} title="Settings" />
          <Divider />
          <Menu.Item onPress={() => {handleLogout()}} title="Logout" />
        </Menu>
      </View>

      <Tab.Navigator
        screenOptions={({ route }) => ({
          tabBarIcon: ({ color, size }) => {
            let iconName: 'home' | 'link' | 'people' = 'home';

            if (route.name === 'Home') {
              iconName = 'home';
            } else if (route.name === 'Triggers') {
              iconName = 'link';
            } else if (route.name === 'Community') {
              iconName = 'people';
            }

            return <MaterialIcons name={iconName} size={size} color={color} />;
          },
          tabBarActiveTintColor: Colors.light.tabIconSelected,
          tabBarInactiveTintColor: Colors.light.tabIconDefault,
          headerShown: false,
        })}
      >
        <Tab.Screen name="Home" component={HomeScreen} />
        <Tab.Screen name="Triggers" component={TriggersScreen} />
        <Tab.Screen name="Community" component={CommunityScreen} />
      </Tab.Navigator>
    </Provider>
  );
}

const styles = StyleSheet.create({
  headerContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: 16,
    backgroundColor: Colors.light.background,
  },
  logo: {
    width: 120,
    height: 40,
    alignSelf: 'center',
  },
});
