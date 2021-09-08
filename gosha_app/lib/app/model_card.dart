import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:gosha_app/controller/app_controller.dart';
import 'package:gosha_app/sdk/model.dart';

class ModelCard extends StatelessWidget {
  var index = 0;

  ModelCard({Key? key, required this.index}) : super(key: key);
  AppController controller = Get.find();

  @override
  Widget build(BuildContext context) {
    return Obx(() {
      var model = controller.models[index];

      return Card(
        elevation: 4,
        color: Colors.grey.shade200,
        child: Column(
          children: [
            Container(
              child: ListTile(
                title: Text(model.Name),
                subtitle: Padding(
                  padding: const EdgeInsets.only(top: 8, bottom: 8),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Chip(
                          label: Text("F"),
                          backgroundColor: model.HttpMethods.Find
                              ? Colors.green.shade300
                              : Colors.grey.shade400),
                      Chip(
                          label: Text("R"),
                          backgroundColor: model.HttpMethods.Read
                              ? Colors.green.shade300
                              : Colors.grey.shade400),
                      Chip(
                          label: Text("C"),
                          backgroundColor: model.HttpMethods.Create
                              ? Colors.lightBlue.shade300
                              : Colors.grey.shade400),
                      Chip(
                          label: Text("U"),
                          backgroundColor: model.HttpMethods.Update
                              ? Colors.yellow.shade300
                              : Colors.grey.shade400),
                      Chip(
                          label: Text("D"),
                          backgroundColor: model.HttpMethods.Delete
                              ? Colors.red.shade300
                              : Colors.grey.shade400),
                      Chip(
                          label: Text("FoC"),
                          backgroundColor: model.HttpMethods.FindOrCreate
                              ? Colors.green.shade100
                              : Colors.grey.shade400),
                      Chip(
                          label: Text("UoC"),
                          backgroundColor: model.HttpMethods.UpdateOrCreate
                              ? Colors.yellow.shade100
                              : Colors.grey.shade400),
                    ],
                  ),
                ),
              ),
              decoration: BoxDecoration(
                  border: Border(bottom: BorderSide(color: Colors.black26))),
            ),
            Padding(
              padding: const EdgeInsets.all(8),
              child: Table(
                defaultVerticalAlignment: TableCellVerticalAlignment.baseline,
                textBaseline: TextBaseline.alphabetic,
                columnWidths: {
                  0: FlexColumnWidth(),
                  1: FlexColumnWidth(),
                  2: FixedColumnWidth(50),
                  3: FixedColumnWidth(50),
                },
                children: [
                  TableRow(
                      decoration: BoxDecoration(
                          border: Border(
                              bottom: BorderSide(color: Colors.black26))),
                      children: [
                        TableCell(child: Padding(
                          padding: const EdgeInsets.only(bottom: 8),
                          child: Text("Name"),
                        )),
                        TableCell(child: Padding(
                          padding: const EdgeInsets.only(bottom: 8),
                          child: Text("Type"),
                        )),
                        TableCell(child: Padding(
                          padding: const EdgeInsets.only(bottom: 8),
                          child: Text("IsDb"),
                        )),
                        TableCell(child: Padding(
                          padding: const EdgeInsets.only(bottom: 8),
                          child: Text("IsType"),
                        )),
                      ]),
                ]..addAll(model.Fields.map((e) => TableRow(children: [
                      Container(
                          padding: EdgeInsets.only(top: 4),
                          child: Text(e.Name)),
                      Container(
                          padding: EdgeInsets.only(top: 4),
                          child: Text(e.Type)),
                      Container(
                          padding: EdgeInsets.only(top: 4),
                          child: Icon(
                            e.IsDbField
                                ? Icons.check_box_outlined
                                : Icons.check_box_outline_blank,
                            size: 18,
                          )),
                      Container(
                          padding: EdgeInsets.only(top: 4),
                          child: Icon(
                            e.IsTypeField
                                ? Icons.check_box_outlined
                                : Icons.check_box_outline_blank,
                            size: 18,
                          )),
                    ]))),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.end,
                children: <Widget>[
                  TextButton(
                    child: const Text('Edit'),
                    onPressed: () {},
                  ),
                  const SizedBox(width: 8),
                ],
              ),
            ),
          ],
        ),
      );
    });
  }
}
