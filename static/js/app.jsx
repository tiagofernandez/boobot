var Boobox = React.createClass({

	getInitialState: function() {
		return { imageUrl: null };
	},

	componentDidMount: function() {
		this.fetchImage();
		setInterval(this.fetchImage, this.props.pollInterval);
	},

	fetchImage: function() {
		$.ajax({
			url: this.props.url,
			dataType: 'json',
			cache: false,
			success: function(data) {
				this.setState({ imageUrl: data.url });
			}.bind(this),
			error: function(xhr, status, err) {
				console.error(this.props.url, status, err.toString());
			}.bind(this)
		});
	},

	render: function() {
		return (
			<div className="boobox">
				<img src={this.state.imageUrl} />
			</div>
		);
	}
});

ReactDOM.render(
	<Boobox url="/api/boobs" pollInterval={5000} />,
	document.getElementById('content')
);
