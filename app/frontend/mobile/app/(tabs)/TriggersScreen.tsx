import React, { useState } from 'react';
import { View, StyleSheet, Modal } from 'react-native';
import { Colors } from '@/constants/Colors';
import FlowChartArea from '@/components/actions/FlowChartArea';
import TechnologySelector from '@/components/actions/TechnologySelector';
import ButtonIcon from '@/components/ButtonIcon';
import { MaterialIcons } from '@expo/vector-icons';

export default function TriggerScreen() {
    const [flow, setFlow] = useState<{ action: string, reactions: string[] }[]>([]);
    const [modalVisible, setModalVisible] = useState<boolean>(false);
    const [selectedActionIndex, setSelectedActionIndex] = useState<number | null>(null);
    const [isAddingAction, setIsAddingAction] = useState<boolean>(false);

    const addAction = (tech: string) => {
        setFlow([...flow, { action: tech, reactions: [] }]);
        setModalVisible(false);
    };

    const addReaction = (tech: string) => {
        if (selectedActionIndex !== null) {
            const newFlow = [...flow];
            newFlow[selectedActionIndex].reactions.push(tech);
            setFlow(newFlow);
            setModalVisible(false);
        }
    };

    const openActionSelector = () => {
        setIsAddingAction(true);
        setModalVisible(true);
    };

    const openReactionSelector = (actionIndex: number) => {
        setIsAddingAction(false);
        setSelectedActionIndex(actionIndex);
        setModalVisible(true);
    };

    const removeAction = (actionIndex: number) => {
        const newFlow = flow.filter((_, index) => index !== actionIndex);
        setFlow(newFlow);
    };

    const saveAction = (flowItem: { action: string, reactions: string[] }) => {
        console.log("Saved action and reactions:", flowItem);
        removeAction(flow.indexOf(flowItem));
    };

    return (
        <View style={styles.container}>
            <View style={styles.addActionContainer}>
                <ButtonIcon
                    title="Add Trigger"
                    onPress={openActionSelector}
                    icon={<MaterialIcons name="add" size={24} color="#FFFFFF" />}
                    backgroundColor={Colors.light.tint}
                    textColor="#FFFFFF"
                />
            </View>

            <FlowChartArea
                flow={flow}
                onAddReaction={openReactionSelector}
                onRemoveAction={removeAction}
                onSaveAction={saveAction}
            />
            <Modal
                animationType="slide"
                transparent={true}
                visible={modalVisible}
                onRequestClose={() => setModalVisible(false)}
            >
                <View style={styles.modalContainer}>
                    <TechnologySelector onTechSelect={isAddingAction ? addAction : addReaction} />
                </View>
            </Modal>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 10,
        backgroundColor: Colors.light.background,
    },
    addActionContainer: {
        marginBottom: 5,
        alignItems: 'flex-start',
    },
    modalContainer: {
        flex: 1,
        justifyContent: 'center',
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
    },
});

// import React, { useState } from 'react';
// import { View, Text, StyleSheet, ScrollView, TouchableOpacity } from 'react-native';
// import { Colors } from '@/constants/Colors';
// // @ts-ignore
// import { Ionicons, FontAwesome5 } from '@expo/vector-icons';
// import { GestureHandlerRootView } from 'react-native-gesture-handler';
// import TechBox from '@/components/actions/TechBox';
// import Draggable from '@/components/actions/Draggable';
// import ActionReactionList from '@/components/actions/ActionReactionList';
// import Video from '@/components/actions/Video';

// interface ActionBox {
//     name: string;
//     icon: React.ReactElement;
// }

// const technologies: ActionBox[] = [
//     { name: 'Google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.google} /> },
//     { name: 'Github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
//     { name: 'Outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.outlook} /> },
//     { name: 'Slack', icon: <FontAwesome5 name="slack" size={30} color={Colors.light.slack} /> },
//     { name: 'Discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.discord} /> },
// ];

// export default function TriggersScreen() {
//     const [service, setService] = useState<ActionBox | undefined>(undefined);

//     const addService = (techName: string) => {
//         const foundTech = technologies.find(tech => tech.name === techName);
//         if (foundTech) {
//             setService(foundTech);
//         }
//     }

//     return (
//         <View style={styles.container}>
//             <View style={styles.servicesContainer}>
//                 <Text style={styles.servicesTitle}>Services</Text>
//                 <ScrollView horizontal showsHorizontalScrollIndicator={false} contentContainerStyle={styles.servicesList}>
//                     {technologies.map((tech, index) => (
//                         <TouchableOpacity key={index} style={styles.techItem} onPress={() => addService(tech.name)}>
//                             <View style={styles.iconContainer}>{tech.icon}</View>
//                             <Text style={styles.techName}>{tech.name}</Text>
//                         </TouchableOpacity>
//                     ))}
//                 </ScrollView>
//             </View>

//             <View style={styles.manageActions}>
//                 <ScrollView showsVerticalScrollIndicator={false}>
//                     <ActionReactionList tech={service}/>
//                 </ScrollView>
//             </View>

//         </View>
//     );
// }

// const styles = StyleSheet.create({
//     container: {
//         flex: 1,
//     },
//     manageActions: {
//         flex: 1,
//         marginTop: 10,
//         padding: 20,
//         borderRadius: 10,
//         backgroundColor: Colors.light.grey,
//         marginHorizontal: 10,
//         height: "100%",
//     },
//     servicesContainer: {
//         marginTop: 10,
//         padding: 20,
//         backgroundColor: Colors.light.grey,
//         borderRadius: 10,
//         marginHorizontal: 10,
//     },
//     servicesTitle: {
//         fontSize: 18,
//         fontWeight: 'bold',
//         marginBottom: 10,
//     },
//     servicesList: {
//         flexDirection: 'row',
//     },
//     techItem: {
//         marginRight: 20,
//         alignItems: 'center',
//         justifyContent: 'center',
//         borderWidth: 1,
//         borderColor: '#ccc',
//         borderRadius: 10,
//         padding: 10,
//         backgroundColor: '#FFFFFF',
//     },
//     iconContainer: {
//         marginBottom: 5,
//     },
//     techName: {
//         fontSize: 16,
//         fontWeight: 'bold',
//     },
// });
