import React from 'react';
import ReactDOM from 'react-dom';
import EventPage from './components/Pages/EventList';
import ContactPage from './components/Pages/ContactList';
import * as serviceWorker from './serviceWorker';
import { BrowserRouter, Switch, Route  } from 'react-router-dom';

ReactDOM.render(
    <BrowserRouter>
        <Switch>
            <Route path="/" exact={true} component={EventPage} />
            <Route path="/contatos" component={ContactPage} />
        </Switch>
    </BrowserRouter>
    , document.getElementById('root'));
serviceWorker.unregister();
