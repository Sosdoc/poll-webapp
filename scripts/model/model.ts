/**
* Answer
*/
export class Answer {

    private text: string;
    private other: boolean;

    constructor(text: string, other: boolean) {
        this.text = text;
        this.other = other;
    }

    public get Text(): string {
        return this.text;
    }

}
    
    

/**
* Poll
*/
export class Poll {
    private title: string;
    private answers: Array<Answer>;

    constructor(title: string, answers: Array<Answer>) {
        this.title = title;
        this.answers = answers;
    }
    
    public get Title(): string {
        return this.title;
    }
    
}
    