import { GroupStackNavigator } from "@/screens/group";
import { SettingsScreen } from "@/screens/settings";
import { ShoppingItemListScreen } from "@/screens/shopping";
import { Colors } from "@/shared/constants";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { Home, Settings, ShoppingBag } from "react-native-feather";

const Tab = createBottomTabNavigator()

type TabBarIconProps = { color: string; size: number }

const tabInfos = [
  {
    name: "Group",
    component: GroupStackNavigator,
    options: {
      title: 'グループ',
      tabBarIcon: ({ color, size }: TabBarIconProps) =>
        <Home stroke={color} width={size} height={size} />
    }
  },
  {
    name: "Shopping",
    component: ShoppingItemListScreen,
    options: {
      title: '買い物',
      tabBarIcon: ({ color, size }: TabBarIconProps) =>
        <ShoppingBag stroke={color} width={size} height={size} />
    }
  },
  {
    name: 'Settings',
    component: SettingsScreen,
    options: {
      title: '設定',
      tabBarIcon: ({ color, size }: TabBarIconProps) =>
        <Settings stroke={color} width={size} height={size} />
    }
  }
]

export const AppNavigator = () => {
  return (
    <Tab.Navigator
      screenOptions={{
        headerShown: false,
        tabBarActiveTintColor: Colors.primary,
        tabBarInactiveTintColor: Colors.mainText,
        tabBarIconStyle: { marginTop: 5 },
        tabBarLabelStyle: { fontSize: 10 },
      }}
    >
      {tabInfos.map((tab) => (
        <Tab.Screen 
          key={tab.name}
          name={tab.name}
          component={tab.component}
          options={tab.options}
        />
      ))}
    </Tab.Navigator>
  )
}