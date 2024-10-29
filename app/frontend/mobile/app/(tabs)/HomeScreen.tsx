import React from 'react';
import { View, StyleSheet, ScrollView, Text, Image } from 'react-native';
import ButtonIcon from '@/components/ButtonIcon';
import Button from '@/components/Button';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';

export default function HomeScreen() {

    return (
        <ScrollView style={styles.container}>
            <PromoItem />
            <Button
                title='Templates'
                onPress={() => console.log('Templates')}
                backgroundColor={Colors.light.tint}
                textColor='#FFFFFF'
                buttonWidth="35%"
                paddingV={7.5}
            />
            <TriggerList />
        </ScrollView>
    );
}

function PromoItem() {
    return (
        <View style={styles.promoContainer}>
            <View style={styles.promoBox}>
                <Text style={styles.promoText}>Try Trigger for 30 days free</Text>
                <ButtonIcon
                    onPress={() => console.log('Start free trial')}
                    title="Start free trial"
                    icon={<MaterialCommunityIcons name="star-shooting" size={24} color={Colors.light.tint} />}
                    backgroundColor="#FFFFFF"
                    textColor={Colors.light.tint}
                />
            </View>
        </View>
    );
}

function TriggerList() {
    const triggers = [
        { id: 1, title: 'Trigger One', imageUrl: require('@/assets/images/image_placeholder.png') },
        { id: 2, title: 'Trigger Two', imageUrl: require('@/assets/images/image_placeholder.png') },
        { id: 3, title: 'Trigger Three', imageUrl: require('@/assets/images/image_placeholder.png') },
    ];

    return (
        <View style={styles.triggerListContainer}>
            <Text style={styles.title}>Your Triggers</Text>
            {triggers.map(trigger => (
                <View key={trigger.id} style={styles.card}>
                    <Image style={styles.cardImage} source={trigger.imageUrl} />
                    <Text style={styles.cardTitle}>{trigger.title}</Text>
                </View>
            ))}
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 16,
    },
    // promo
    promoContainer: {
        alignItems: 'center',
        marginBottom: 10,
    },
    promoBox: {
        width: '100%',
        backgroundColor: Colors.light.tint,
        padding: 10,
        borderRadius: 10,
        alignItems: 'center',
    },
    promoText: {
        color: '#fff',
        fontSize: 18,
        marginBottom: 10,
        textAlign: 'center',
    },
    promoButton: {
        backgroundColor: '#fff',
        paddingVertical: 10,
        paddingHorizontal: 20,
        borderRadius: 5,
    },
    promoButtonText: {
        color: Colors.light.tint,
        fontWeight: 'bold',
    },
    // trigger list
    triggerListContainer: {
        marginTop: 10,
        marginBottom: 30,
    },
    title: {
        fontSize: 22,
        fontWeight: 'bold',
        marginBottom: 16,
    },
    card: {
        backgroundColor: '#fff',
        borderRadius: 10,
        marginBottom: 20,
        padding: 16,
        shadowColor: '#000',
        shadowOpacity: 0.1,
        shadowRadius: 5,
        shadowOffset: { width: 0, height: 2 },
        alignItems: 'center',
    },
    cardImage: {
        width: 300,
        height: 200,
        marginBottom: 10,
        borderRadius: 10,
    },
    cardTitle: {
        fontSize: 16,
        fontWeight: 'bold',
    },
});

// import React from 'react';
// import { View, Text, Image, TouchableOpacity, StyleSheet } from 'react-native';
// import { MaterialIcons } from '@expo/vector-icons';
// import { Menu, Divider, Button, Provider } from 'react-native-paper';

// export default function HomeScreen() {
//     const [visible, setVisible] = React.useState(false);

//     const openMenu = () => setVisible(true);
//     const closeMenu = () => setVisible(false);

//     return (
//         <Provider>
//             <View style={styles.container}>
//                 <View style={styles.header}>
//                     <Text style={styles.greeting}>Hello!</Text>
//                     <Menu
//                         visible={visible}
//                         onDismiss={closeMenu}
//                         anchor={
//                             <TouchableOpacity onPress={openMenu}>
//                                 <Image
//                                     source={{ uri: 'user-profile-image-url' }}
//                                     style={styles.profile}
//                                 />
//                             </TouchableOpacity>
//                         }>
//                         <Menu.Item onPress={() => { }} title="Settings" />
//                         <Divider />
//                         <Menu.Item onPress={() => { }} title="Log Out" />
//                     </Menu>
//                 </View>
//             </View>
//         </Provider>
//     );
// }

// const styles = StyleSheet.create({
//     container: {
//         flex: 1,
//         padding: 16,
//     },
//     header: {
//         flexDirection: 'row',
//         justifyContent: 'space-between',
//         alignItems: 'center',
//     },
//     greeting: {
//         fontSize: 24,
//     },
//     profile: {
//         width: 40,
//         height: 40,
//         borderRadius: 20,
//     },
// });
