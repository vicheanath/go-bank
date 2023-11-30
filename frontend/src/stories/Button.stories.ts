
import type { Meta, StoryObj } from '@storybook/react';
import { Button as ButtonComponents } from '../ui/Button';
import SvgOutlineGlobe from '../icons/OutlineGlobe';
import React from 'react';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories#default-export
const meta = {
  title: 'UI/Button',
  component: ButtonComponents,
  tags: ['autodocs'],
  argTypes: { onClick: { action: "clicked" } },
} satisfies Meta<typeof ButtonComponents>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Button: Story = {
  args: {
    size: 'big',
    children: 'Button',
    color: 'secondary',
    icon: null,
  },
};

export const ButtonWithIcon: Story = {
  args: {
    size: 'big',
    children: 'Button',
    color: 'secondary',
    icon: React.createElement(SvgOutlineGlobe, { className: 'w-3 h-3' }),
  },
};

