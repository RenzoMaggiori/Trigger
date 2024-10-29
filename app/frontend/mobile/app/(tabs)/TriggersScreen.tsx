import React, { useState } from 'react';
import { View, StyleSheet, Modal, TouchableOpacity, Text } from 'react-native';
import { Colors } from '@/constants/Colors';
import FlowChartArea from '@/components/actions/FlowChartArea';
import ProviderSelector from '@/components/actions/ProviderSelector';
import ActionSelector from '@/components/actions/ActionsSelector';
import ButtonIcon from '@/components/ButtonIcon';
import { MaterialIcons } from '@expo/vector-icons';

export default function TriggerScreen() {
    const [flow, setFlow] = useState<{ provider: string, action: string, reactions: { provider: string, reaction: string }[] }[]>([]);
    const [modalVisible, setModalVisible] = useState<boolean>(false);
    const [selectedActionIndex, setSelectedActionIndex] = useState<number | null>(null);
    const [isAddingAction, setIsAddingAction] = useState<boolean>(false);
    const [selectedProvider, setSelectedProvider] = useState<string | null>(null);
    const [showActionSelector, setShowActionSelector] = useState<boolean>(false);

    const addAction = (action: string) => {
        if (selectedProvider) {
            setFlow([...flow, { provider: selectedProvider, action, reactions: [] }]);
            closeModal();
        }
    };

    const addReaction = (reaction: string) => {
        if (selectedActionIndex !== null && selectedProvider) {
            setFlow(prevFlow => {
                const updatedFlow = [...prevFlow];
                updatedFlow[selectedActionIndex].reactions.push({ provider: selectedProvider, reaction });
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

    const saveTrigger = (flowItem: { provider: string, action: string, reactions: { provider: string, reaction: string }[] }) => {
        console.log("Saved Trigger:", flowItem);
        removeAction(flow.indexOf(flowItem));
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
