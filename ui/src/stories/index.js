import React from 'react';
import '@storybook/addon-actions/register';
import '@storybook/addon-knobs/register';
import { storiesOf } from '@storybook/react';
import { withKnobs, text, number, select, boolean } from '@storybook/addon-knobs';

//Atoms
import Snackbar from '../components/Atoms/Snackbar';

//Templates
import EventListTemplate from '../components/Templates/EventList';
import ContactListTemplate from '../components/Templates/ContactList';

//Pages
import ContactListPage from '../components/Pages/ContactList';
import EventListPage from '../components/Pages/EventList';

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
    .add('EventListTemplate', () => <EventListTemplate />)
    .add('ContactListTemplate', () => <ContactListTemplate />);

storiesOf('Pages', module)
    .add('ContactListPage', () => <ContactListPage />)
    .add('EventListPage', () => <EventListPage />);
