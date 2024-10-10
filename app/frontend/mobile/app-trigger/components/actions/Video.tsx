import { Colors } from "@/constants/Colors";
import React, { useRef, useState } from "react";
import { Text, TouchableOpacity, StyleSheet } from "react-native";
import DraggableFlatList, { OpacityDecorator, ScaleDecorator, ShadowDecorator, useOnCellActiveAnimation } from "react-native-draggable-flatlist";

import { GestureHandlerRootView } from "react-native-gesture-handler";
import Animated from "react-native-reanimated";

interface Data {
    key: string;
    label: string;
}

export default function Video() {
    const ref = useRef(null);
    
    const [data, setData] = useState<Data[]>([
        { key: "a", label: "a" },
        { key: "b", label: "b" },
        { key: "c", label: "c" },
        { key: "d", label: "d" },
        { key: "e", label: "e" },
        { key: "f", label: "f" },
        { key: "g", label: "g" },
    ]);


    const renderItem = ({ item, drag }: { item: Data, drag: () => void }) => {
        const { isActive } = useOnCellActiveAnimation();
        return (
            <ScaleDecorator>
                <OpacityDecorator activeOpacity={0.6}>
                    <ShadowDecorator>
                        <TouchableOpacity
                            style={[
                                styles.rowItem,
                                {
                                    height: 100,
                                    justifyContent: "center",
                                    alignItems: "center",
                                    backgroundColor: "grey",
                                    elevation: isActive ? 30 : 0,
                                }
                            ]}
                            activeOpacity={1}
                            onLongPress={drag}>
                                <Animated.View>
                                    <Text>{item.label}</Text>
                                </Animated.View>
                        </TouchableOpacity>
                    </ShadowDecorator>
                </OpacityDecorator>
            </ScaleDecorator>
        );
    };

    console.log(data);


    return (
        <GestureHandlerRootView>
            <Text>Video</Text>
            <DraggableFlatList
                ref={ref}
                data={data}
                keyExtractor={ (item: Data) => item.key }
                onDragEnd={ ({data}) => setData(data) }
                renderItem={renderItem}
            >
            </DraggableFlatList>
        </GestureHandlerRootView>
    );
}

const styles = StyleSheet.create({
    manageActions: {
        flex: 1,
        marginTop: 10,
        padding: 20,
        borderRadius: 10,
        backgroundColor: Colors.light.grey,
        marginHorizontal: 10,
        height: "100%",
    },
    rowItem: {
        marginVertical: 10,
        marginHorizontal: 20,
        borderRadius: 10,
    },
});

