import React, { useState } from 'react';
import { View, StyleSheet, Modal, TouchableOpacity, Text } from 'react-native';
import { Colors } from '@/constants/Colors';
import FlowChartArea from '@/components/actions/FlowChartArea';
import ProviderSelector from '@/components/actions/ProviderSelector';
import ActionSelector from '@/components/actions/ActionsSelector';
import ButtonIcon from '@/components/ButtonIcon';
import { MaterialIcons } from '@expo/vector-icons';
import { TriggersService } from '@/api/triggers/service';

function transformFlowItemToApiFormat(flowItem: { action: { id: string, name: string }, reactions: { id: string, name: string }[] }, inputValues: any) {
    const nodeNames = [];
    const nodes = [];

    // Trigger node
    nodes.push({
        node_id: "action1",
        action_id: flowItem.action.id,
        name: flowItem.action.name,
        input: inputValues,
        output: {},
        parents: [],
        children: flowItem.reactions.length > 0 ? ["action2"] : [],
        x_pos: 10,
        y_pos: 10,
    });
    nodeNames.push(flowItem.action.name);

    // Reaction nodes
    flowItem.reactions.forEach((reaction, index) => {
        const nodeIndex = index + 2;
        nodes.push({
            node_id: `action${nodeIndex}`,
            action_id: reaction.id,
            name: reaction.name,
            input: inputValues,
            output: {},
            parents: [`action${nodeIndex - 1}`],
            children: index < flowItem.reactions.length - 1 ? [`action${nodeIndex + 1}`] : [],
            x_pos: 10,
            y_pos: 10 * nodeIndex,
        });
        nodeNames.push(reaction.name);
    });

    const name = nodeNames.join(' > ');
    return { name, nodes };
}

export default function TriggerScreen() {
    const [flow, setFlow] = useState<{ provider: string; action: { id: string; name: string }; reactions: { provider: string; id: string; name: string }[] }[]>([]);
    const [modalVisible, setModalVisible] = useState(false);
    const [selectedActionIndex, setSelectedActionIndex] = useState<number | null>(null);
    const [isAddingAction, setIsAddingAction] = useState(false);
    const [selectedProvider, setSelectedProvider] = useState<string | null>(null);
    const [showActionSelector, setShowActionSelector] = useState(false);

    const addAction = (action: { id: string; name: string }) => {
        if (selectedProvider) {
            setFlow([...flow, { provider: selectedProvider, action, reactions: [] }]);
            closeModal();
        }
    };

    const addReaction = (reaction: { id: string; name: string }) => {
        if (selectedActionIndex !== null && selectedProvider) {
            setFlow(prevFlow => {
                const updatedFlow = [...prevFlow];
                updatedFlow[selectedActionIndex].reactions.push({ provider: selectedProvider, ...reaction });
                return updatedFlow;
            });
            closeModal();
        }
    };

    const openActionSelector = () => {
        setIsAddingAction(true);
        setSelectedProvider(null);
        setModalVisible(true);
    };

    const openReactionSelector = (actionIndex: number) => {
        setIsAddingAction(false);
        setSelectedActionIndex(actionIndex);
        setSelectedProvider(null);
        setModalVisible(true);
    };

    const selectProvider = (provider: string) => {
        setSelectedProvider(provider);
        setShowActionSelector(true);
    };

    const removeAction = (actionIndex: number) => {
        setFlow(prevFlow => prevFlow.filter((_, index) => index !== actionIndex));
    };

    const saveTrigger = (actionIndex: number, inputValues: any) => {
        const flowItem = flow[actionIndex];
        const formattedData = transformFlowItemToApiFormat(flowItem, inputValues);
        TriggersService.addTrigger(formattedData);
        removeAction(actionIndex);
    };

    const closeModal = () => {
        setModalVisible(false);
        setShowActionSelector(false);
        setSelectedProvider(null);
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
                onSaveTrigger={saveTrigger}
            />
            <Modal
                animationType="slide"
                transparent={true}
                visible={modalVisible}
                onRequestClose={closeModal}
            >
                <View style={styles.modalContainer}>
                    <View style={styles.modalContent}>
                        <TouchableOpacity style={styles.closeButton} onPress={closeModal}>
                            <Text style={styles.closeButtonText}>×</Text>
                        </TouchableOpacity>
                        {!showActionSelector ? (
                            <ProviderSelector onProviderSelect={selectProvider} />
                        ) : (
                            <ActionSelector
                                provider={selectedProvider}
                                onActionSelect={isAddingAction ? addAction : addReaction}
                                type={isAddingAction ? 'trigger' : 'reaction'}
                            />
                        )}
                    </View>
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
    modalContent: {
        backgroundColor: '#fff',
        borderRadius: 10,
        marginHorizontal: 30,
        position: 'relative',
    },
    closeButton: {
        position: 'absolute',
        top: 10,
        right: 10,
        zIndex: 1,
    },
    closeButtonText: {
        fontSize: 24,
        color: '#000',
    },
});
