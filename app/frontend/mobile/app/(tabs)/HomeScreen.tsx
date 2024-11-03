import React, { useEffect, useState, useCallback } from 'react';
import { View, StyleSheet, ScrollView, Text, TouchableOpacity, Modal } from 'react-native';
import ButtonIcon from '@/components/ButtonIcon';
import Button from '@/components/Button';
import { MaterialCommunityIcons, MaterialIcons } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';
import { TriggersService } from '@/api/triggers/service';
import { useFocusEffect } from '@react-navigation/native';

export default function HomeScreen() {
    const [workspaces, setWorkspaces] = useState<Workspace[]>([]);

    useFocusEffect(
        useCallback(() => {
            const loadWorkspaces = async () => {
                try {
                    const fetchedWorkspaces = await TriggersService.getTriggers();
                    // console.log('fetchedWorkspaces:', fetchedWorkspaces);

                    const updatedWorkspaces = await Promise.all(
                        fetchedWorkspaces.map(async (workspace: Workspace) => {
                            const nodesWithDetails = await Promise.all(
                                workspace.nodes.map(async (node) => {
                                    const actionDetails = await TriggersService.getActionById(node.action_id);
                                    return {
                                        ...node,
                                        actionDetails,
                                    };
                                })
                            );
                            return {
                                ...workspace,
                                nodes: nodesWithDetails,
                            };
                        })
                    );

                    setWorkspaces(updatedWorkspaces);
                } catch (error) {
                    console.error('Failed to load triggers:', error);
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
            <TriggerList workspaces={workspaces} setWorkspaces={setWorkspaces} />
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
    name: string;
    nodes: {
        node_id: string;
        action_id: string;
        status: string;
        actionDetails: {
            id: string;
            input: string[];
            output: string[];
            provider: string;
            type: string;
            action: string;
        };
    }[];
}

function TriggerList({ workspaces, setWorkspaces }: { workspaces: Workspace[], setWorkspaces: React.Dispatch<React.SetStateAction<Workspace[]>> }) {
    const [modalVisible, setModalVisible] = useState(false);
    const [selectedWorkspaceId, setSelectedWorkspaceId] = useState<string | null>(null);

    const handleTriggerAction = async (workspaceId: string, isActive: boolean) => {
        try {
            if (isActive) {
                await TriggersService.stopTrigger(workspaceId);
                console.log(`Stopped trigger for workspace ID: ${workspaceId}`);
            } else {
                await TriggersService.startTrigger(workspaceId);
                console.log(`Started trigger for workspace ID: ${workspaceId}`);
            }
            setWorkspaces(prev => prev.map(workspace => {
                if (workspace.id === workspaceId) {
                    return {
                        ...workspace,
                        nodes: workspace.nodes.map(node => ({
                            ...node,
                            status: isActive ? 'inactive' : 'active',
                        })),
                    };
                }
                return workspace;
            }));
        } catch (error) {
            console.error("Error handling trigger action:", error);
        }
    };

    const handleDeleteWorkspace = async () => {
        if (selectedWorkspaceId) {
            try {
                await TriggersService.deleteTrigger(selectedWorkspaceId);
                setWorkspaces(prev => prev.filter(workspace => workspace.id !== selectedWorkspaceId));
                console.log(`Deleted trigger with ID: ${selectedWorkspaceId}`);
            } catch (error) {
                console.error("Error deleting trigger:", error);
            } finally {
                setModalVisible(false);
            }
        }
    };

    const openDeleteModal = (workspaceId: string) => {
        setSelectedWorkspaceId(workspaceId);
        setModalVisible(true);
    };

    return (
        <View style={styles.triggerListContainer}>
            <Text style={styles.title}>Your Triggers</Text>
            {workspaces.length > 0 ? (
                workspaces.map((workspace) => (
                    <View key={workspace.id} style={styles.workspaceCard}>
                        {/* <Text style={styles.workspaceTitle}>Trigger ID: {workspace.id}</Text> */}
                        <Text style={styles.workspaceTitle}>{workspace.name}</Text>
                        <View style={styles.nodesContainer}>
                            {workspace.nodes.map((node) => {
                                const isTrigger = node.actionDetails?.type === 'trigger';

                                return (
                                    <View
                                        key={node.node_id}
                                        style={[
                                            styles.nodeCard,
                                            isTrigger && styles.triggerNodeCard,
                                        ]}
                                    >
                                        <View style={styles.nodeHeader}>
                                            <View style={styles.nodeDetails}>
                                                <Text
                                                    style={[
                                                        styles.nodeText,
                                                        isTrigger && styles.whiteText,
                                                    ]}
                                                >
                                                    Provider: {node.actionDetails?.provider}
                                                </Text>
                                                <Text
                                                    style={[
                                                        styles.nodeText,
                                                        isTrigger && styles.whiteText,
                                                    ]}
                                                >
                                                    {isTrigger ? 'Action' : 'Reaction'}: {node.actionDetails?.action}
                                                </Text>
                                                <View style={styles.actionDetailsContainer}>
                                                    <Text
                                                        style={[
                                                            styles.actionDetailText,
                                                            isTrigger && styles.whiteText,
                                                        ]}
                                                    >
                                                        Inputs: {node.actionDetails?.input.join(', ')}
                                                    </Text>
                                                    <Text
                                                        style={[
                                                            styles.actionDetailText,
                                                            isTrigger && styles.whiteText,
                                                        ]}
                                                    >
                                                        Outputs: {node.actionDetails?.output.join(', ')}
                                                    </Text>
                                                </View>
                                            </View>
                                            <View style={styles.rightSection}>
                                                <Text
                                                    style={[
                                                        styles.nodeStatus,
                                                        isTrigger && styles.whiteText,
                                                    ]}
                                                >
                                                    {node.status === 'active' ? 'Active' : 'Inactive'}
                                                </Text>
                                            </View>
                                        </View>
                                    </View>
                                );
                            })}
                        </View>
                        <View style={styles.buttonRow}>
                            <Button
                                title='Delete Trigger'
                                onPress={() => openDeleteModal(workspace.id)}
                                backgroundColor='#FFFFFF'
                                textColor={Colors.light.tintDark}
                                paddingV={10}
                                buttonWidth="45%"
                            />
                            <Button
                                title={workspace.nodes.some(node => node.status === 'active') ? 'Stop Trigger' : 'Start Trigger'}
                                onPress={() => handleTriggerAction(workspace.id, workspace.nodes.some(node => node.status === 'active'))}
                                backgroundColor={Colors.light.tint}
                                textColor='#FFFFFF'
                                paddingV={10}
                                buttonWidth="45%"
                            />
                        </View>
                    </View>
                ))
            ) : (
                <Text style={styles.noTriggersText}>No Triggers available.</Text>
            )}

            <Modal
                animationType="slide"
                transparent={true}
                visible={modalVisible}
                onRequestClose={() => setModalVisible(false)}
            >
                <View style={styles.modalContainer}>
                    <View style={styles.modalContent}>
                        <Text style={styles.modalTitle}>Delete Trigger</Text>
                        <Text style={styles.modalMessage}>
                            Are you sure you want to delete this trigger?
                        </Text>
                        <View style={styles.modalButtons}>
                            <TouchableOpacity
                                onPress={() => setModalVisible(false)}
                                style={styles.cancelButton}
                            >
                                <Text style={styles.buttonText}>CANCEL</Text>
                            </TouchableOpacity>
                            <TouchableOpacity
                                onPress={handleDeleteWorkspace}
                                style={styles.acceptButton}
                            >
                                <Text style={styles.buttonText}>DELETE</Text>
                            </TouchableOpacity>
                        </View>
                    </View>
                </View>
            </Modal>
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
    workspaceCard: {
        backgroundColor: Colors.light.tintLight,
        borderRadius: 10,
        marginBottom: 20,
        padding: 16,
        shadowColor: '#000',
        shadowOpacity: 0.1,
        shadowRadius: 5,
        shadowOffset: { width: 0, height: 2 },
    },
    workspaceTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 10,
    },
    nodesContainer: {
        marginBottom: 10,
    },
    nodeCard: {
        backgroundColor: '#f9f9f9',
        borderRadius: 8,
        padding: 10,
        marginBottom: 10,
        shadowColor: '#000',
        shadowOpacity: 0.05,
        shadowRadius: 3,
        shadowOffset: { width: 0, height: 1 },
    },
    triggerNodeCard: {
        backgroundColor: Colors.light.tintDark,
    },
    nodeHeader: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'flex-start',
    },
    nodeDetails: {
        flex: 1,
    },
    rightSection: {
        flexDirection: 'column',
        alignItems: 'flex-end',
    },
    nodeText: {
        fontSize: 14,
        fontWeight: 'bold',
        color: Colors.light.tintDark,
    },
    whiteText: {
        color: '#fff',
    },
    actionDetailsContainer: {
        marginTop: 5,
    },
    actionDetailText: {
        fontSize: 12,
    },
    nodeStatus: {
        fontSize: 14,
        fontWeight: 'bold',
        color: '#666',
        marginBottom: 5,
    },
    actionButton: {
        flexDirection: 'row',
        alignItems: 'center',
        backgroundColor: Colors.light.tint,
        paddingVertical: 4,
        paddingHorizontal: 8,
        borderRadius: 20,
        marginTop: 5,
    },
    invertedActionButton: {
        backgroundColor: '#fff',
    },
    actionButtonTxt: {
        color: '#fff',
        marginLeft: 5,
    },
    invertedActionButtonTxt: {
        color: Colors.light.tintDark,
    },
    noTriggersText: {
        textAlign: 'center',
        color: '#888',
    },
    deleteButton: {
        backgroundColor: '#ff4d4d',
        padding: 10,
        borderRadius: 5,
        alignItems: 'center',
        marginTop: 10,
    },
    deleteButtonText: {
        color: '#fff',
        fontWeight: 'bold',
    },
    modalContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'rgba(0, 0, 0, 0.7)',
    },
    modalContent: {
        width: 300,
        backgroundColor: '#fff',
        padding: 20,
        borderRadius: 10,
        alignItems: 'center',
    },
    modalTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 15,
    },
    modalMessage: {
        fontSize: 16,
        marginBottom: 20,
    },
    modalButtons: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        width: '100%',
    },
    cancelButton: {
        backgroundColor: '#d3d3d3',
        padding: 10,
        borderRadius: 5,
        width: '45%',
        alignItems: 'center',
    },
    acceptButton: {
        backgroundColor: Colors.light.tint,
        padding: 10,
        borderRadius: 5,
        width: '45%',
        alignItems: 'center',
    },
    buttonText: {
        color: '#fff',
        fontWeight: 'bold',
    },
    buttonRow: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        marginTop: 10,
    },
});
