import React, { useState } from 'react';
import { View, Image, TouchableOpacity, StyleSheet, Text, Dimensions, GestureResponderEvent } from 'react-native';
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
  const [triggersButtonColor, setTriggersButtonColor] = useState(Colors.light.tint);

  const openMenu = () => setMenuVisible(true);
  const closeMenu = () => setMenuVisible(false);

  const router = useRouter();

  const handleLogout = () => {
    router.push('/LandingPage' as const);
  };

  const handleSettings = () => {
    closeMenu();
    router.push('/Settings' as const);
  };

  const handleTriggersPress = (onPress: (e: GestureResponderEvent) => void) => (e: GestureResponderEvent) => {
    setTriggersButtonColor(Colors.light.tintDark);
    onPress(e);
  };

  const handleTabPress = (routeName: string) => {
    if (routeName !== 'Triggers') {
      setTriggersButtonColor(Colors.light.tint);
    }
  };

  const data = {
    logo: require('../../assets/images/logo.png'),
  };

  return (
    <Provider>
      <View style={styles.navbar}>
        <View style={styles.logoContainer}>
          <Image source={data.logo} style={styles.logo} />
        </View>
        <View style={styles.menuContainer}>
          <Menu
            visible={menuVisible}
            onDismiss={closeMenu}
            anchor={
              <TouchableOpacity onPress={openMenu}>
                <MaterialIcons name="account-circle" size={32} color={Colors.light.tint} />
              </TouchableOpacity>
            }
          >
            <View style={styles.sideMenu}>
              <Menu.Item
                onPress={handleSettings}
                title="Settings"
                leadingIcon={() => <MaterialIcons name="settings" size={24} color={Colors.light.tint} />}
              />
              <Divider />
              <Menu.Item
                onPress={handleLogout}
                title="Logout"
                leadingIcon={() => <MaterialIcons name="logout" size={24} color={Colors.light.tint} />}
              />
            </View>
          </Menu>
        </View>
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
        screenListeners={({ route }) => ({
          tabPress: () => handleTabPress(route.name),
        })}
      >
        <Tab.Screen name="Home" component={HomeScreen} />

        <Tab.Screen
          name="Triggers"
          component={TriggersScreen}
          options={{
            tabBarButton: (props) => (
              <TouchableOpacity
                {...props}
                style={[styles.triggersButton, { backgroundColor: triggersButtonColor }]}
                onPress={props.onPress && handleTriggersPress(props.onPress)}
              >
                <MaterialIcons name="add" size={30} color="#fff" />
                <Text style={styles.triggersText}>Triggers</Text>
              </TouchableOpacity>
            ),
          }}
        />

        <Tab.Screen name="Community" component={CommunityScreen} />
      </Tab.Navigator>
    </Provider>
  );
}

const styles = StyleSheet.create({
  navbar: {
    height: 80,
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    backgroundColor: '#fff',
    paddingTop: 20,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
    paddingHorizontal: 20,
  },
  logoContainer: {
    flex: 1,
    alignItems: 'center',
  },
  menuContainer: {
    top: 35,
    position: 'absolute',
    right: 20,
  },
  sideMenu: {
    marginTop: 20,
  },
  logo: {
    resizeMode: 'contain',
    height: 30,
  },
  triggersButton: {
    width: 70,
    height: 70,
    borderRadius: 35,
    justifyContent: 'center',
    alignItems: 'center',
    position: 'absolute',
    bottom: 10,
    left: (Dimensions.get('window').width / 2) - 35,
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 3,
    },
    shadowOpacity: 0.4,
    shadowRadius: 4,
    elevation: 6,
  },
  triggersText: {
    color: '#fff',
    fontSize: 10,
  },
});
