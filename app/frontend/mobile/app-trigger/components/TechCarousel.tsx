import React, { useRef, useEffect } from 'react';
import { View, Text, StyleSheet, FlatList, Animated, Easing } from 'react-native';
import { Colors } from '@/constants/Colors';

interface Technology {
    name: string;
    icon: JSX.Element;
}

interface TechCarouselProps {
    technologies: Technology[];
}

const TechCarousel: React.FC<TechCarouselProps> = ({ technologies }) => {
    const flatListRef = useRef<FlatList>(null);
    const scrollX = useRef(new Animated.Value(0)).current;
    const currentOffset = useRef(0);

    useEffect(() => {
        const listenerId = scrollX.addListener(({ value }) => {
            currentOffset.current = value;
        });

        const scrollInterval = setInterval(() => {
            const nextOffset = currentOffset.current + 75;

            Animated.timing(scrollX, {
                toValue: nextOffset,
                duration: 1000,
                easing: Easing.linear,
                useNativeDriver: false,
            }).start(() => {
                if (currentOffset.current >= technologies.length * 100) {
                    flatListRef.current?.scrollToOffset({
                        offset: 0,
                        animated: false,
                    });
                    scrollX.setValue(0);
                } else {
                    flatListRef.current?.scrollToOffset({
                        offset: nextOffset,
                        animated: true,
                    });
                }
            });
        }, 100);
        return () => {
            clearInterval(scrollInterval);
            scrollX.removeListener(listenerId);
        };
    }, [scrollX, technologies.length]);

    return (
        <View style={styles.carouselContainer}>
            <FlatList
                ref={flatListRef}
                data={technologies.concat(technologies)}
                horizontal
                showsHorizontalScrollIndicator={false}
                keyExtractor={(item, index) => index.toString()}
                renderItem={({ item }) => (
                    <View style={styles.techItem}>
                        {item.icon}
                        <Text style={styles.techName}>{item.name}</Text>
                    </View>
                )}
                scrollEnabled={false}
            />
        </View>
    );
};

const styles = StyleSheet.create({
    carouselContainer: {
        marginVertical: 20,
    },
    techItem: {
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'flex-start',
        marginHorizontal: 20,
    },
    techName: {
        fontSize: 16,
        fontWeight: 'bold',
        color: Colors.light.tintLight,
        marginLeft: 10,
    },
});

export default TechCarousel;

//------------- NO ANIMATION VERSION

// import React from 'react';
// import { View, Text, StyleSheet, FlatList } from 'react-native';

// interface Technology {
//     name: string;
//     icon: JSX.Element;
// }

// interface TechCarouselProps {
//     technologies: Technology[];
// }

// const TechCarousel: React.FC<TechCarouselProps> = ({ technologies }) => {
//     return (
//         <View style={styles.carouselContainer}>
//             <FlatList
//                 data={technologies}
//                 horizontal
//                 showsHorizontalScrollIndicator={false}
//                 keyExtractor={(item) => item.name}
//                 renderItem={({ item }) => (
//                     <View style={styles.techItem}>
//                         {item.icon}
//                         <Text style={styles.techName}>{item.name}</Text>
//                     </View>
//                 )}
//             />
//         </View>
//     );
// };

// const styles = StyleSheet.create({
//     carouselContainer: {
//         marginVertical: 20,
//     },
//     techItem: {
//         flexDirection: 'row',
//         alignItems: 'center',
//         justifyContent: 'flex-start',
//         marginHorizontal: 20,
//     },
//     techName: {
//         fontSize: 16,
//         fontWeight: 'bold',
//         color: '#888',
//         marginLeft: 10,
//     },
// });

// export default TechCarousel;
