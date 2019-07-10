import React from 'react';
import '@storybook/addon-actions/register';
import '@storybook/addon-knobs/register';
import { storiesOf } from '@storybook/react';
import { withKnobs, text, number, select, boolean } from '@storybook/addon-knobs';

//Atoms
import Snackbar from '../components/Atoms/Snackbar';

//Templates
import EventList from '../components/Templates/EventList';
import ContactList from '../components/Templates/ContactList';

//Pages
/*TODO*/

storiesOf('Atoms', module)
    .addDecorator(withKnobs)
    .add('Snackbar', () => <Snackbar
        message={text('Some message', 'Hey something happens')}
        open={boolean("Open:", false)}
        variant={select('status', {
            Success: 'success',
            Warning: 'warning',
            Error: 'error',
            Info: 'info',
        }, 'info')}
    />);

storiesOf('Templates', module)
    .add('EventList', () => <EventList />)
    .add('ContactList', () => <ContactList name="test"/>);

storiesOf('Pages', module)
    .add('EventList', () => <EventList />)
    .add('ContactList', () => <ContactList name="test"/>);
