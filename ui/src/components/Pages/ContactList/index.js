import React from 'react';
import Template from '../../Templates/ContactList';
import Snackbar from '../../Atoms/Snackbar';
import Styled from 'styled-components';
import { Box, Container } from '@material-ui/core';
import { create,
    getAll,
    update,
    clear
} from './api';

const Content = Styled(Box)`
	margin-top: 5%;
    padding: 5;
`

class ContactList extends React.Component {

    constructor(props) {
		super(props)
        this.state = {
            message: '',
            status: 'info',
            snack: false,
        }
    }

    showSnack = (message, status) => {
        this.setState({
            ...this.state,
            message: message,
            status: status,
            snack: true
        });
		setTimeout(() => {
			this.setState({ ...this.state, snack: false });
		}, 6000);
    }

	onRowAdd = newData => create(newData)
		.then(response => this.showSnack(response, 'success'))
		.catch(error => {
			this.showSnack(error.toString(), 'error')
			throw error
		})

	onRowUpdate = (newData, oldData) => update(newData)
		.then(response => this.showSnack(response, 'success'))
		.catch(error => {
			this.showSnack(error.toString(), 'error')
			throw error
		})

	onRowDelete = oldData => clear(oldData.phone)
		.then(response => this.showSnack(response, 'success'))
		.catch(error => {
			this.showSnack(error.toString(), 'error')
			throw error
		})

    render() {
        const { message, snack, status, data } = this.state
        return (
			<Content>
				<Snackbar
					message={message}
					open={snack}
					variant={status}
				/>
				<Container maxWidth="lg">
					<Template
						onRowAdd={newData => this.onRowAdd(newData)}
						onRowUpdate={this.onRowUpdate}
						onRowDelete={this.onRowDelete}
                        getInit={getAll}
					/>
				</Container>
			</Content>
        )
    }
}

export default ContactList
