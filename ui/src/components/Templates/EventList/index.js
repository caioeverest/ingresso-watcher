import React, { forwardRef } from 'react';
import MaterialTable from 'material-table';
import { AddBox,
    ArrowUpward,
    Check,
    ChevronLeft,
    ChevronRight,
    Clear,
    DeleteOutline,
    Edit,
    FilterList,
    FirstPage,
    LastPage,
    Remove,
    SaveAlt,
    Search,
    ViewColumn
} from "@material-ui/icons";

const tableIcons = {
    Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
    Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
    Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
    DetailPanel: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
    Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
    Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
    FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
    LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
    NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    PreviousPage: forwardRef((props, ref) => <ChevronLeft {...props} ref={ref} />),
    ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
    SortArrow: forwardRef((props, ref) => <ArrowUpward {...props} ref={ref} />),
    ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref} />),
    ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />)
};

class EventList extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            columns: [
                { title: 'Id', field: 'id' },
                { title: 'Nome', field: 'name' },
            ],
            data: [
                {id: '12344', name: 'test'},
            ],
        }
    }

    componentDidMount() {
       console.log("test")
    }

    onRowAdd = newData => new Promise(resolve => {
        setTimeout(() => {
            resolve();
            const { data } = this.state;
            data.push(newData);
            this.setState({ ...this.state, data });
        }, 600);
    })

    onRowUpdate = (newData, oldData) => new Promise(resolve => {
        setTimeout(() => {
            resolve();
            const { data } = this.state;
            data[data.indexOf(oldData)] = newData;
            this.setState({ ...this.state, data });
        }, 600);
    })

    onRowDelete = oldData => new Promise(resolve => {
        setTimeout(() => {
            resolve();
            const { data } = this.state;
            data.splice(data.indexOf(oldData), 1);
            this.setState({ ...this.state, data });
        }, 600);
    })

    render() {
        const { data, columns } = this.state
        return (
            <MaterialTable
                title="Lista de eventos"
                icons={tableIcons}
                columns={columns}
                data={data}
                editable={{
                    onRowAdd: this.onRowAdd,
                    onRowUpdate: this.onRowUpdate,
                    onRowDelete: this.onRowDelete,
                }}
            />
        );
    }
}

export default EventList
