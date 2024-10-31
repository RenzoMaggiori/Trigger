import React, { useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity, ScrollView } from 'react-native';
import { MaterialIcons, Entypo } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';

interface FlowChartAreaProps {
    flow: {
        provider: string,
        action: { id: string, name: string },
        reactions: { id: string, provider: string, name: string }[]
    }[];
    onAddReaction: (actionIndex: number) => void;
    onRemoveAction: (actionIndex: number) => void;
    onSaveTrigger: (actionIndex: number) => void;
}

export default function FlowChartArea({ flow, onAddReaction, onRemoveAction, onSaveTrigger }: FlowChartAreaProps) {
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
                            <Text style={styles.actionText}>Provider: {flowItem.provider}</Text>
                            <Text style={styles.actionText}>Action: {flowItem.action.name}</Text>
                        </View>
                    </TouchableOpacity>

                    {flowItem.reactions.map((reaction, reactionIndex) => (
                        <TouchableOpacity key={reactionIndex} onPress={() => handleSelectReaction(actionIndex, reactionIndex)}>
                            <View style={styles.reactionContainer}>
                                <Text style={styles.reactionText}>Provider: {reaction.provider}</Text>
                                <Text style={styles.reactionText}>Reaction: {reaction.name}</Text>
                            </View>
                        </TouchableOpacity>
                    ))}

                    {selectedItem.actionIndex === actionIndex && selectedItem.type === 'action' && (
                        <View style={styles.infoCard}>
                            <Text style={styles.infoText}>Provider: {flowItem.provider}</Text>
                            <Text style={styles.infoText}>Action: {flowItem.action.name}</Text>
                        </View>
                    )}

                    {selectedItem.actionIndex === actionIndex && selectedItem.type === 'reaction' && selectedItem.reactionIndex !== null && (
                        <View style={styles.infoCard}>
                            <Text style={styles.infoText}>Reaction Selected</Text>
                            <Text>{flowItem.reactions[selectedItem.reactionIndex!].provider}: {flowItem.reactions[selectedItem.reactionIndex!].name}</Text>
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
                            <TouchableOpacity style={styles.actionButton} onPress={() => onSaveTrigger(actionIndex)}>
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
        backgroundColor: Colors.light.tintDark,
        padding: 10,
        borderRadius: 5,
        marginBottom: 10,
    },
    actionText: {
        fontWeight: 'bold',
        color: '#fff',
    },
    reactionContainer: {
        backgroundColor: '#fff',
        padding: 10,
        borderRadius: 5,
        marginTop: 5,
    },
    reactionText: {
        fontWeight: 'bold',
        color: Colors.light.tintDark,
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
