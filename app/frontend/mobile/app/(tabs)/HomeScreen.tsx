import React from 'react';
import { View, Text, Image, TouchableOpacity, StyleSheet } from 'react-native';
import { MaterialIcons } from '@expo/vector-icons';
import { Menu, Divider, Button, Provider } from 'react-native-paper';

export default function HomeScreen() {
    const [visible, setVisible] = React.useState(false);

    const openMenu = () => setVisible(true);
    const closeMenu = () => setVisible(false);

    return (
        <Provider>
            <View style={styles.container}>
                <View style={styles.header}>
                    <Text style={styles.greeting}>Hello!</Text>
                    <Menu
                        visible={visible}
                        onDismiss={closeMenu}
                        anchor={
                            <TouchableOpacity onPress={openMenu}>
                                <Image
                                    source={{ uri: 'user-profile-image-url' }}
                                    style={styles.profile}
                                />
                            </TouchableOpacity>
                        }>
                        <Menu.Item onPress={() => { }} title="Settings" />
                        <Divider />
                        <Menu.Item onPress={() => { }} title="Log Out" />
                    </Menu>
                </View>
            </View>
        </Provider>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 16,
    },
    header: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
    },
    greeting: {
        fontSize: 24,
    },
    profile: {
        width: 40,
        height: 40,
        borderRadius: 20,
    },
});
