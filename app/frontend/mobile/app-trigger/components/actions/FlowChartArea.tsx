import React from 'react';
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
    return (
        <ScrollView style={styles.container}>
            {flow.map((flowItem, index) => (
                <View key={index} style={styles.flowItem}>
                    <View style={styles.actionContainer}>
                        <Text style={styles.actionText}>Action: {flowItem.action}</Text>
                    </View>

                    {flowItem.reactions.map((reaction, reactionIndex) => (
                        <View key={reactionIndex} style={styles.reactionContainer}>
                            <Text style={styles.reactionText}>Reaction: {reaction}</Text>
                        </View>
                    ))}

                    <View style={styles.buttonsContainer}>
                        <TouchableOpacity style={styles.actionButton} onPress={() => onAddReaction(index)}>
                            <MaterialIcons name="add" size={24} color="#fff" />
                            <Text style={styles.actionButtonTxt}>Add Reaction</Text>
                        </TouchableOpacity>
                        <View style={styles.options}>
                            <TouchableOpacity style={styles.actionButton} onPress={() => onRemoveAction(index)}>
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
        marginTop: 20,
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
});
