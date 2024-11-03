import React, { useEffect, useState, useCallback } from 'react';
import { View, StyleSheet, ScrollView, Text, Image } from 'react-native';
import ButtonIcon from '@/components/ButtonIcon';
import Button from '@/components/Button';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';
import { TriggersService } from '@/api/triggers/service';
import { useFocusEffect } from '@react-navigation/native';

export default function HomeScreen() {
    const [workspaces, setWorkspaces] = useState([]);

    useFocusEffect(
        useCallback(() => {
            const loadWorkspaces = async () => {
                try {
                    const fetchedWorkspaces = await TriggersService.getTriggers();
                    console.log('fetchedWorkspaces:', fetchedWorkspaces);
                    setWorkspaces(fetchedWorkspaces);
                } catch (error) {
                    console.error('Failed to load workspaces:', error);
                }
            };

            loadWorkspaces();
        }, [])
    );

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
            <TriggerList workspaces={workspaces} />
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

interface Workspace {
    id: string;
    user_id: string;
    nodes: {
        node_id: string;
        action_id: string;
        status: string;
    }[];
}

function TriggerList({ workspaces }: { workspaces: Workspace[] }) {
    return (
        <View style={styles.triggerListContainer}>
            <Text style={styles.title}>Your Workspaces</Text>
            {workspaces.length > 0 ? (
                workspaces.map(workspace => (
                    <View key={workspace.id} style={styles.card}>
                        <Text style={styles.cardTitle}>Workspace ID: {workspace.id}</Text>
                        <Text style={styles.cardSubtitle}>User ID: {workspace.user_id}</Text>
                        <View style={styles.nodesContainer}>
                            {workspace.nodes.map((node, index) => (
                                <View key={node.node_id} style={styles.nodeItem}>
                                    <Text style={styles.nodeText}>
                                        {index + 1}. {node.node_id} - {node.action_id} (Status: {node.status})
                                    </Text>
                                </View>
                            ))}
                        </View>
                    </View>
                ))
            ) : (
                <Text style={styles.noTriggersText}>No workspaces available.</Text>
            )}
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 16,
    },
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
    triggerListContainer: {
        marginTop: 10,
        marginBottom: 30,
    },
    title: {
        fontSize: 22,
        fontWeight: 'bold',
        marginBottom: 16,
        textAlign: 'center',
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
    },
    cardTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 5,
    },
    cardSubtitle: {
        fontSize: 14,
        color: '#666',
        marginBottom: 10,
    },
    nodesContainer: {
        marginTop: 10,
    },
    nodeItem: {
        marginBottom: 5,
    },
    nodeText: {
        fontSize: 14,
        color: '#333',
    },
    noTriggersText: {
        textAlign: 'center',
        fontSize: 16,
        color: '#888',
    },
});

// import React from 'react';
// import { View, StyleSheet, ScrollView, Text, Image } from 'react-native';
// import ButtonIcon from '@/components/ButtonIcon';
// import Button from '@/components/Button';
// import { MaterialCommunityIcons } from '@expo/vector-icons';
// import { Colors } from '@/constants/Colors';
// import { TriggersService } from '@/api/triggers/service';

// export default function HomeScreen() {

//     return (
//         <ScrollView style={styles.container}>
//             <PromoItem />
//             <Button
//                 title='Templates'
//                 onPress={() => console.log('Templates')}
//                 backgroundColor={Colors.light.tint}
//                 textColor='#FFFFFF'
//                 buttonWidth="35%"
//                 paddingV={7.5}
//             />
//             <TriggerList />
//         </ScrollView>
//     );
// }

// function PromoItem() {
//     return (
//         <View style={styles.promoContainer}>
//             <View style={styles.promoBox}>
//                 <Text style={styles.promoText}>Try Trigger for 30 days free</Text>
//                 <ButtonIcon
//                     onPress={() => console.log('Start free trial')}
//                     title="Start free trial"
//                     icon={<MaterialCommunityIcons name="star-shooting" size={24} color={Colors.light.tint} />}
//                     backgroundColor="#FFFFFF"
//                     textColor={Colors.light.tint}
//                 />
//             </View>
//         </View>
//     );
// }

// function TriggerList() {
//     const triggers = [
//         { id: 1, title: 'Trigger One', imageUrl: require('@/assets/images/image_placeholder.png') },
//         { id: 2, title: 'Trigger Two', imageUrl: require('@/assets/images/image_placeholder.png') },
//         { id: 3, title: 'Trigger Three', imageUrl: require('@/assets/images/image_placeholder.png') },
//     ];

//     console.log('triggers:', TriggersService.getTriggers());

//     return (
//         <View style={styles.triggerListContainer}>
//             <Text style={styles.title}>Your Triggers</Text>
//             {triggers.map(trigger => (
//                 <View key={trigger.id} style={styles.card}>
//                     <Image style={styles.cardImage} source={trigger.imageUrl} />
//                     <Text style={styles.cardTitle}>{trigger.title}</Text>
//                 </View>
//             ))}
//         </View>
//     );
// }

// const styles = StyleSheet.create({
//     container: {
//         flex: 1,
//         padding: 16,
//     },
//     // promo
//     promoContainer: {
//         alignItems: 'center',
//         marginBottom: 10,
//     },
//     promoBox: {
//         width: '100%',
//         backgroundColor: Colors.light.tint,
//         padding: 10,
//         borderRadius: 10,
//         alignItems: 'center',
//     },
//     promoText: {
//         color: '#fff',
//         fontSize: 18,
//         marginBottom: 10,
//         textAlign: 'center',
//     },
//     promoButton: {
//         backgroundColor: '#fff',
//         paddingVertical: 10,
//         paddingHorizontal: 20,
//         borderRadius: 5,
//     },
//     promoButtonText: {
//         color: Colors.light.tint,
//         fontWeight: 'bold',
//     },
//     // trigger list
//     triggerListContainer: {
//         marginTop: 10,
//         marginBottom: 30,
//     },
//     title: {
//         fontSize: 22,
//         fontWeight: 'bold',
//         marginBottom: 16,
//     },
//     card: {
//         backgroundColor: '#fff',
//         borderRadius: 10,
//         marginBottom: 20,
//         padding: 16,
//         shadowColor: '#000',
//         shadowOpacity: 0.1,
//         shadowRadius: 5,
//         shadowOffset: { width: 0, height: 2 },
//         alignItems: 'center',
//     },
//     cardImage: {
//         width: 300,
//         height: 200,
//         marginBottom: 10,
//         borderRadius: 10,
//     },
//     cardTitle: {
//         fontSize: 16,
//         fontWeight: 'bold',
//     },
// });
