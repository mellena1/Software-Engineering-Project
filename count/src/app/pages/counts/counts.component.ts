import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';

import {Count} from '../../data_models/count';
import {CountService} from '../../services/count.service';

@Component({
  selector: 'app-counts',
  templateUrl: './counts.component.html',
  styleUrls: ['./counts.component.css']
})
export class CountsComponent implements OnInit {
  constructor(private countService: CountService) {}
  ngOnInit() {}
}
