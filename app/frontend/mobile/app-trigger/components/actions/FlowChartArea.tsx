import React, { useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity, ScrollView } from 'react-native';
import { MaterialIcons, Entypo } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';

interface FlowChartAreaProps {
    flow: { action: string, reactions: string[] }[];
    onAddReaction: (actionIndex: number) => void;
    onRemoveAction: (actionIndex: number) => void;
    onSaveAction: (flowItem: { action: string, reactions: string[] }) => void;
}

export default function FlowChartArea({ flow, onAddReaction, onRemoveAction, onSaveAction }: FlowChartAreaProps) {
    const [selectedItem, setSelectedItem] = useState<{ type: 'action' | 'reaction', actionIndex: number | null, reactionIndex?: number | null }>({
        type: 'action',
        actionIndex: null,
        reactionIndex: null,
    });

    const handleSelectAction = (actionIndex: number) => {
        setSelectedItem(prevState => ({
            type: 'action',
            actionIndex: prevState.actionIndex === actionIndex ? null : actionIndex,
            reactionIndex: null,
        }));
    };

    const handleSelectReaction = (actionIndex: number, reactionIndex: number) => {
        setSelectedItem(prevState => ({
            type: 'reaction',
            actionIndex,
            reactionIndex: prevState.reactionIndex === reactionIndex ? null : reactionIndex,
        }));
    };

    return (
        <ScrollView style={styles.container}>
            {flow.map((flowItem, actionIndex) => (
                <View key={actionIndex} style={styles.flowItem}>

                    <TouchableOpacity onPress={() => handleSelectAction(actionIndex)}>
                        <View style={styles.actionContainer}>
                            <Text style={styles.actionText}>Action: {flowItem.action}</Text>
                        </View>
                    </TouchableOpacity>

                    {flowItem.reactions.map((reaction, reactionIndex) => (
                        <TouchableOpacity key={reactionIndex} onPress={() => handleSelectReaction(actionIndex, reactionIndex)}>
                            <View style={styles.reactionContainer}>
                                <Text style={styles.reactionText}>Reaction: {reaction}</Text>
                            </View>
                        </TouchableOpacity>
                    ))}

                    {selectedItem.actionIndex === actionIndex && selectedItem.type === 'action' && (
                        <View style={styles.infoCard}>
                            <Text style={styles.infoText}>Action: {flowItem.action}</Text>
                            <Text>Here goes Action Data</Text>
                        </View>
                    )}

                    {selectedItem.actionIndex === actionIndex && selectedItem.type === 'reaction' && selectedItem.reactionIndex !== null && (
                        <View style={styles.infoCard}>
                            <Text style={styles.infoText}>Reaction: {flowItem.reactions[selectedItem.reactionIndex ?? 0]}</Text>
                            <Text>Here goes Reaction Data</Text>
                        </View>
                    )}

                    <View style={styles.buttonsContainer}>
                        <TouchableOpacity style={styles.actionButton} onPress={() => onAddReaction(actionIndex)}>
                            <MaterialIcons name="add" size={24} color="#fff" />
                            <Text style={styles.actionButtonTxt}>Add Reaction</Text>
                        </TouchableOpacity>
                        <View style={styles.options}>
                            <TouchableOpacity style={styles.actionButton} onPress={() => onRemoveAction(actionIndex)}>
                                <Entypo name="cross" size={24} color="#fff" />
                                <Text style={styles.actionButtonTxt}>Remove</Text>
                            </TouchableOpacity>
                            <TouchableOpacity style={styles.actionButton} onPress={() => onSaveAction(flowItem)}>
                                <MaterialIcons name="save-alt" size={24} color="#fff" />
                                <Text style={styles.actionButtonTxt}>Save</Text>
                            </TouchableOpacity>
                        </View>
                    </View>
                </View>
            ))}
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
    },
    flowItem: {
        marginBottom: 20,
        backgroundColor: Colors.light.tintLight,
        padding: 10,
        borderRadius: 10,
        elevation: 3,
    },
    actionContainer: {
        backgroundColor: '#fff',
        padding: 10,
        borderRadius: 5,
        marginBottom: 10,
    },
    actionText: {
        fontWeight: 'bold',
    },
    reactionContainer: {
        backgroundColor: '#ffeaa7',
        padding: 10,
        borderRadius: 5,
        marginTop: 5,
    },
    reactionText: {
        fontWeight: 'bold',
    },
    buttonsContainer: {
        flexDirection: 'column',
        marginTop: 10,
    },
    options: {
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    actionButton: {
        flexDirection: 'row',
        alignItems: 'center',
        backgroundColor: Colors.light.tint,
        padding: 10,
        borderRadius: 20,
        marginTop: 7.5,
    },
    actionButtonTxt: {
        color: '#fff',
        marginLeft: 5,
    },
    infoCard: {
        backgroundColor: '#f5f5f5',
        padding: 15,
        borderRadius: 8,
        marginTop: 30,
    },
    infoText: {
        fontWeight: 'bold',
        color: Colors.light.tint,
    },
});
