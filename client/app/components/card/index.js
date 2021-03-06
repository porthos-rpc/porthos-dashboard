import React from 'react';
import {AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip} from 'Recharts';
import './styles.css'

class Card extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        var historyChart;

        if (this.props.data.history.length > 1) {
            historyChart = <AreaChart width={280} height={80} data={this.props.data.history} margin={{top: 0, right: 0, left: 0, bottom: 0}}>
                <Tooltip/>
                <Area type='monotone' dataKey='throughput' stroke='#2c4c30' fill='#c0d0b6' />
                <Area type='monotone' dataKey='responseTime' stroke='#7b7815' fill='#cab114' />
            </AreaChart>
        } else {
            historyChart = <div className="notEnoughData">Not enough data...</div>
        }

        return (
            <div className="card">
                <h3>{this.props.data.serviceName}</h3> 
                <h2>{this.props.data.methodName}</h2>
                <div className="chartAvgs">
                    <div className="chart">
                        {historyChart} 
                    </div>
                    <div className="avgs">
                        <div className="avgResponseTime">{this.props.data.avgResponseTime} ms</div> 
                        <div className="avgThroughput">{this.props.data.avgThroughput} RPM</div> 
                    </div>
                </div>
                <div className="cardCommonMetric">
                    <div className="left">ResponseTime</div> <div className="right"><b>Min: {this.props.data.minResponseTime} ms Max: {this.props.data.maxResponseTime} ms</b></div>
                </div>
                <div className="cardCommonMetric">
                    <div className="left">Throughput</div> <div className="right"><b>Min: {this.props.data.minThroughput} RPM Max: {this.props.data.maxThroughput} RPM</b></div>
                </div>
            </div>
        );
    }

}

export default Card;

