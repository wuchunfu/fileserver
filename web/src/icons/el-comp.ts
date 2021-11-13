/**
 * element-plus icon: https://element-plus.gitee.io/zh-CN/component/icon.html
 * @description 图标库按需注册
 * @author wcf
 * @example <el-icon :size="20"> <edit /> </el-icon>
 */

// el-icon
import {
  ElAlert,
  ElAside,
  ElBacktop,
  ElButton,
  ElCheckbox,
  ElCol,
  ElContainer,
  ElDivider,
  ElDropdown,
  ElDropdownItem,
  ElDropdownMenu,
  ElFooter,
  ElForm,
  ElFormItem,
  ElHeader,
  ElIcon,
  ElInput,
  ElMain,
  ElMenu,
  ElMenuItem,
  ElMenuItemGroup,
  ElRadio,
  ElRadioButton,
  ElRadioGroup,
  ElRow,
  ElScrollbar,
  ElSelect,
  ElSlider,
  ElSubMenu,
  ElTable,
  ElTableColumn,
  ElTabPane,
  ElTabs,
  ElTag,
  ElTooltip,
} from 'element-plus';

// 所需的组件
export const components = [
  ElAlert,
  ElAside,
  ElButton,
  ElSelect,
  ElRow,
  ElCol,
  ElForm,
  ElFormItem,
  ElInput,
  ElTabs,
  ElTabPane,
  ElCheckbox,
  ElIcon,
  ElDivider,
  ElBacktop,
  ElDropdown,
  ElDropdownMenu,
  ElDropdownItem,
  ElContainer,
  ElHeader,
  ElSlider,
  ElMain,
  ElFooter,
  ElMenu,
  ElMenuItem,
  ElMenuItemGroup,
  ElSubMenu,
  ElRadio,
  ElRadioButton,
  ElRadioGroup,
  ElTooltip,
  ElScrollbar,
  ElTableColumn,
  ElTag,
  ElTable,
];

// 注册
export default (app: any) => {
  components.forEach((component) => {
    app.component(component.name, component);
  });
};
