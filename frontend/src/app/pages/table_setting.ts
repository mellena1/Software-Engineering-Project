export class TableSetting {
  settings: object;

  constructor(columns: object) {
    this.settings = {
      actions: {
        position: "right"
      },
      add: {
        confirmCreate: true,
        addButtonContent:
          '<div class="btn btn-success btn-sm mx-1 px-2">New Room</div>',
        createButtonContent:
          '<div class="btn btn-success btn-sm mx-1 px-2">Add</div>',
        cancelButtonContent:
          '<div></div>',
        //   '<div class="btn btn-danger btn-sm mx-1 px-2">Cancel</div>'
      },
      edit: {
        confirmSave: true,
        editButtonContent:
          '<div class="btn btn-primary btn-sm mx-1 px-2">Edit</div>',
        saveButtonContent:
          '<div class="btn btn-success btn-sm mx-1 px-2">Save</div>',
        cancelButtonContent:
          '<div class="btn btn-danger btn-sm mx-1 px-2">Cancel</div>'
      },
      delete: {
        confirmDelete: true,
        deleteButtonContent:
          '<div class="btn btn-danger btn-sm mx-1 px-2">Delete</div>'
      },
      attr: {
        class: "table table-bordered table-striped test"
      },
      hideSubHeader: true,
      columns: columns
    };
  }
}
