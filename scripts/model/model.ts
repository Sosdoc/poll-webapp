namespace Polls {
	/**
	* Poll
	*/
	class Poll {
		private title: string;
		private answers: Array<Answer>;

		constructor(title: string, answers: Array<Answer>) {
			this.title = title;
			this.answers = answers;
		}
	}
	
	/**
 	* Answer
 	*/
	class Answer {

		private text: string;
		private other: boolean;

		constructor(text: string, other: boolean) {
			this.text = text;
			this.other = other;
		}
	}
}