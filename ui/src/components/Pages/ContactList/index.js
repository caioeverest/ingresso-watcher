import React from 'react'
import Template from '../../Templates/ContactList'
import Styled from 'styled-components'
import { Box } from '@material-ui/core';
import { create,
    getAll,
    getByPhone,
    update,
    delete
} from './api'

const Content = Styled(Box)`
    padding: 5;
`

class ContactList extends React.Component {

    constructor() {
        this.state = {
            errorMessage: false,
            data: [],
        }
    }

    componentDidMount() {
        getAll().then(data => this.state({
            ...this.state,
            data
        }))
    }

    onRowAdd = newData => create(newData)
        .then(() => {
            const { data } = this.state;
            data.push(newData);
            this.setState({ ...this.state, data });
        })

    onRowUpdate = (newData, oldData) => this.props.onRowUpdate(newData, oldData)
        .then(() => {
            const { data } = this.state;
            data[data.indexOf(oldData)] = newData;
            this.setState({ ...this.state, data });
        })

    onRowDelete = oldData => this.props.onRowDelete(oldData)
        .then(() => {
            const { data } = this.state;
            data.splice(data.indexOf(oldData), 1);
            this.setState({ ...this.state, data });
        })

    render() {
        const { errorMessage, data } = this.state
        return (
            <Content>
                <Template
                    onRowAdd={ this.onRowAdd }
                    onRowUpdate={ this.onRowUpdate }
                    onRowDelete={ this.onRowDelete }
                    data= { data }
                />
            </Content>
        )
    }
}

export default ContactList
