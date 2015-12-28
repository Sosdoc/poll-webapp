import * as model from "../model/model"

interface PollState {
    
}

export class PollComponent extends React.Component<model.Poll, PollState> {
    render() {
        return <div> 
        <h1>{this.props.Title}</h1> 
        <ul>
            <li>Answer number 1</li>
        </ul>        
        </div>
    }
}