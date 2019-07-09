import React from 'react';
import '@storybook/addon-actions/register';
import '@storybook/addon-knobs/register';

import { storiesOf } from '@storybook/react';
import { withKnobs, text, number, select, boolean } from '@storybook/addon-knobs';

import Snacks from '../components/Atoms/Snacks';

import EventList from '../components/Templates/EventList';
import ContactList from '../components/Templates/ContactList';

storiesOf('Atoms', module)
    .addDecorator(withKnobs)
    .add('Snacks', () => <Snacks
        message={text('Some message', 'Hey something happens')}
        open={boolean('it is open', false)}
        timeout={number('Time that it will be shown um ms', 5)}
        status={select('status', {
            Success: 'success',
            Warning: 'warning',
            Error: 'error',
            Info: 'info',
        }, 'info')}
    />);

storiesOf('Templates', module)
    .add('EventList', () => <EventList />)
    .add('ContactList', () => <ContactList name="test"/>);
